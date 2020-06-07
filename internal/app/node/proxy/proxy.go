package proxy

import (
	"sync"

	"github.com/google/gopacket/layers"
	"github.com/spf13/viper"
	"mmesh.dev/mmesh/internal/pkg/ipnet"
)

type ServiceType int

const (
	ServiceTypeNone ServiceType = iota
	ServiceTypeKubernetes
	ServiceTypeProxy64
)

type proxyNamespaceType int

const (
	proxyNamespaceTypeNone proxyNamespaceType = iota
	proxyNamespaceTypeKubernetes
	proxyNamespaceTypeProxy64
)

const NamespaceNone string = "-"

type proxyNsName string
type proxySvcName string
type proxyPortName string

type proxyPort struct {
	name    string
	proto   layers.IPProtocol
	port    int32
	wg      sync.WaitGroup
	active  bool
	running bool
	closeCh chan struct{}
}

type proxySvc struct {
	name  string
	af    ipnet.AddressFamily
	ip    string
	vip   string
	ports map[proxyPortName]*proxyPort
}

type proxyNamespace struct {
	nsType proxyNamespaceType
	svc    map[proxySvcName]*proxySvc
}

type proxyServiceMap struct {
	ns map[proxyNsName]*proxyNamespace
	sync.RWMutex
}

var proxySvcMap *proxyServiceMap

func newProxyServiceMap() *proxyServiceMap {
	return &proxyServiceMap{
		ns: make(map[proxyNsName]*proxyNamespace),
	}
}

func newProxyNamespace(proxyNsType proxyNamespaceType) *proxyNamespace {
	return &proxyNamespace{
		nsType: proxyNsType,
		svc:    make(map[proxySvcName]*proxySvc),
	}
}

func newProxySvc(svcName, svcIP, vIP string, af ipnet.AddressFamily) *proxySvc {
	return &proxySvc{
		name:  svcName,
		af:    af,
		ip:    svcIP,
		vip:   vIP,
		ports: make(map[proxyPortName]*proxyPort),
	}
}

func newProxyPort(portName string, proto layers.IPProtocol, port int32) *proxyPort {
	return &proxyPort{
		name:    portName,
		proto:   proto,
		port:    port,
		active:  false,
		running: false,
		closeCh: make(chan struct{}),
	}
}

func (pxsm *proxyServiceMap) setNS(ns string, proxyNsType proxyNamespaceType) {
	pxsm.Lock()
	defer pxsm.Unlock()

	if _, ok := pxsm.ns[proxyNsName(ns)]; !ok {
		pxsm.ns[proxyNsName(ns)] = newProxyNamespace(proxyNsType)
	}
}

func (pxsm *proxyServiceMap) setSvc(proxyNsType proxyNamespaceType, namespace, svcName, svcIP, vIP string, af ipnet.AddressFamily) {
	pxsm.setNS(namespace, proxyNsType)

	pxsm.Lock()
	defer pxsm.Unlock()

	ns := pxsm.ns[proxyNsName(namespace)]

	if svc, ok := ns.svc[proxySvcName(svcName)]; !ok {
		ns.svc[proxySvcName(svcName)] = newProxySvc(svcName, svcIP, vIP, af)
	} else {
		if svc.ip != svcIP || svc.vip != vIP {
			ns.svc[proxySvcName(svcName)] = newProxySvc(svcName, svcIP, vIP, af)
		}
	}
}

func (pxsm *proxyServiceMap) setPort(proxyNsType proxyNamespaceType, namespace, svcName, svcIP, vIP, portName string, proto layers.IPProtocol, port int32, af ipnet.AddressFamily) {
	pxsm.setSvc(proxyNsType, namespace, svcName, svcIP, vIP, af)

	pxsm.deletePort(namespace, svcName, portName)

	pxsm.Lock()
	defer pxsm.Unlock()

	ns := pxsm.ns[proxyNsName(namespace)]
	svc := ns.svc[proxySvcName(svcName)]

	svc.ports[proxyPortName(portName)] = newProxyPort(portName, proto, port)
}

func (pxsm *proxyServiceMap) deletePort(namespace, svcName, portName string) {
	pxsm.Lock()
	defer pxsm.Unlock()

	if ns, ok := pxsm.ns[proxyNsName(namespace)]; ok {
		if svc, ok := ns.svc[proxySvcName(svcName)]; ok {
			if port, ok := svc.ports[proxyPortName(portName)]; ok {
				if port.running {
					port.closeCh <- struct{}{}
					port.wg.Wait()
				}
			}
			delete(svc.ports, proxyPortName(portName))
		}
	}
}

func (pxsm *proxyServiceMap) runningPort(namespace, svcName, portName string) bool {
	pxsm.Lock()
	defer pxsm.Unlock()

	if ns, ok := pxsm.ns[proxyNsName(namespace)]; ok {
		if svc, ok := ns.svc[proxySvcName(svcName)]; ok {
			if port, ok := svc.ports[proxyPortName(portName)]; ok {
				if port.running {
					return true
				}
			}
		}
	}

	return false
}

func (pxsm *proxyServiceMap) fwdSvc(namespace, svcName string) {
	agentPort := int32(viper.GetInt("agent.port"))

	pxsm.Lock()
	defer pxsm.Unlock()

	if ns, ok := pxsm.ns[proxyNsName(namespace)]; ok {
		if svc, ok := ns.svc[proxySvcName(svcName)]; ok {
			for _, port := range svc.ports {
				if port.proto != layers.IPProtocolTCP || port.port == agentPort {
					continue
				}
				if !port.running {
					port.wg.Add(1)
					port.running = true
					go port.fwdCtl(namespace, svcName, svc.vip, svc.ip, svc.af)
				}
			}
		}
	}
}

func SetPort(proxyServiceType ServiceType, namespace, svcName, svcIP, vIP, portName string, proto layers.IPProtocol, port int32, af ipnet.AddressFamily) {
	var pxNsType proxyNamespaceType

	switch proxyServiceType {
	case ServiceTypeNone:
		pxNsType = proxyNamespaceTypeNone
	case ServiceTypeKubernetes:
		pxNsType = proxyNamespaceTypeKubernetes
	case ServiceTypeProxy64:
		pxNsType = proxyNamespaceTypeProxy64
	default:
		pxNsType = proxyNamespaceTypeNone
	}

	if proxySvcMap == nil {
		proxySvcMap = newProxyServiceMap()
	}

	proxySvcMap.setPort(pxNsType, namespace, svcName, svcIP, vIP, portName, proto, port, af)
}

func DeletePort(namespace, svcName, portName string) {
	if proxySvcMap == nil {
		proxySvcMap = newProxyServiceMap()
		return
	}

	proxySvcMap.deletePort(namespace, svcName, portName)
}

func FwdSvc(namespace, svcName string) {
	if proxySvcMap == nil {
		proxySvcMap = newProxyServiceMap()
		return
	}

	proxySvcMap.fwdSvc(namespace, svcName)
}

func RunningPort(namespace, svcName, portName string) bool {
	if proxySvcMap == nil {
		return false
	}

	return proxySvcMap.runningPort(namespace, svcName, portName)
}
