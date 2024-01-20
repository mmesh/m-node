package svcs

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

// const dnsPortAlt int = 53535

type dnsName string

type handler struct {
	nxnc rpc.NetworkAPIClient
}

type dnsMap struct {
	rr map[dnsName]*dnsRR
	sync.RWMutex
}

type dnsRR struct {
	dns.RR
	timestamp int64
}

var dnsCache *dnsMap

// dnsAgent method implementation of NxNetwork gRPC Service
func DNSAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	dnsPort := mnet.LocalNode().DNSPort()

	srv := &dns.Server{
		Addr: ":" + strconv.Itoa(dnsPort),
		Net:  "udp",
		Handler: &handler{
			nxnc: w.NxNC,
		},
		ReusePort: true,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			xlog.Alertf("Unable to set up DNS agent: %v", err)
		}
	}()

	<-w.QuitChan

	if err := srv.Shutdown(); err != nil {
		xlog.Errorf("Unable to shutdown DNS listener: %v", err)
	}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

func (h *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)

	msg.SetReply(r)
	msg.Authoritative = true
	name := msg.Question[0].Name

	if rr := queryDNSCache(name, r.Question[0].Qtype); rr != nil {
		msg.Answer = append(msg.Answer, rr)

		if err := w.WriteMsg(msg); err != nil {
			xlog.Errorf("DNS: %v", err)
		}
		return
	}

	var ipv4, ipv6 string

	// resolver for mmesh endpoints:
	//    endpointName.mmesh.local.
	//    endpointName.mm.*
	//    endpointName.namespace.mmesh.local.
	//    endpointName.namespace.mm.*

	xlog.Debugf("[dns] Received DNS query: %s", name)

	s := strings.Split(name, ".")

	var dnsName string

	if len(s) == 4 {
		if (s[1] == "mmesh" && s[2] == "local") ||
			s[1] == "mm" ||
			s[1] == "mmesh" {
			dnsName = s[0]
		}
	}

	if len(s) == 5 {
		if (s[2] == "mmesh" && s[3] == "local") ||
			s[2] == "mm" ||
			s[2] == "mmesh" {
			dnsName = fmt.Sprintf("%s.%s", s[0], s[1])
		}
	}

	if len(dnsName) > 0 {
		ipv4, ipv6 = mnet.LocalNode().Router().RIB().DNSQuery(dnsName)

		xlog.Debugf("[dns] Name: %s | IPv4: %s | IPv6: %s", name, ipv4, ipv6)
	}

	switch r.Question[0].Qtype {
	case dns.TypeA:
		addrs := make([]string, 0)

		if len(ipv4) > 0 {
			addrs = append(addrs, ipv4)
		}

		if len(addrs) == 0 {
			if !strings.Contains(name, "mmesh") {
				addrs = localResolver(name)
			}
		}

		for _, addr := range addrs {
			rr := &dns.A{
				Hdr: dns.RR_Header{
					Name:   name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    300,
				},
				A: net.ParseIP(addr),
			}
			msg.Answer = append(msg.Answer, rr)
			go updateDNSCache(rr)
		}
	case dns.TypeAAAA:
		addrs := make([]string, 0)

		if len(ipv6) > 0 {
			addrs = append(addrs, ipv6)
		}

		if len(addrs) == 0 {
			addrs = mmesh64Resolver(name)
		}

		for _, addr := range addrs {
			rr := &dns.AAAA{
				Hdr: dns.RR_Header{
					Name:   name,
					Rrtype: dns.TypeAAAA,
					Class:  dns.ClassINET,
					Ttl:    300,
				},
				AAAA: net.ParseIP(addr),
			}
			msg.Answer = append(msg.Answer, rr)
			go updateDNSCache(rr)
		}
	}

	// if len(msg.Answer) == 0 {
	// 	msg.Authoritative = false

	// 	if in, err := dnsProxy(r.Question); err != nil {
	// 		xlog.Errorf("DNS: %v", err)
	// 	} else {
	// 		msg.Answer = in.Answer
	// 	}
	// }

	if err := w.WriteMsg(msg); err != nil {
		xlog.Errorf("DNS: %v", err)
	}
}

func newDNSMap() *dnsMap {
	return &dnsMap{
		rr: make(map[dnsName]*dnsRR),
	}
}

func newDNSRR(rr dns.RR) *dnsRR {
	return &dnsRR{
		RR:        rr,
		timestamp: time.Now().Unix(),
	}
}

func (dm *dnsMap) set(rr dns.RR) {
	dm.Lock()
	defer dm.Unlock()

	switch r := rr.(type) {
	case *dns.A:
		xlog.Debugf("DNS: Caching A RR for name %s (%s)", r.Hdr.Name, r.A.String())
		dm.rr[dnsName(r.Hdr.Name)] = newDNSRR(r)
	case *dns.AAAA:
		xlog.Debugf("DNS: Caching AAAA RR for name %s (%s)", r.Hdr.Name, r.AAAA.String())
		dm.rr[dnsName(r.Hdr.Name)] = newDNSRR(r)
	}
}

func (dm *dnsMap) get(name string, qtype uint16) dns.RR {
	dm.Lock()
	defer dm.Unlock()

	dnsRR, ok := dm.rr[dnsName(name)]
	if !ok {
		return nil
	}

	if dnsRR.Header().Rrtype != qtype {
		return nil
	}

	if time.Now().Unix()-dnsRR.timestamp > int64(dnsRR.Header().Ttl) {
		// ttl expired
		xlog.Debugf("DNS: TTL expired for name %s", name)
		delete(dm.rr, dnsName(name))
		return nil
	}

	return dnsRR.RR
}

func updateDNSCache(rr dns.RR) {
	if dnsCache == nil {
		dnsCache = newDNSMap()
	}

	dnsCache.set(rr)
}

func queryDNSCache(name string, qtype uint16) dns.RR {
	if dnsCache == nil {
		dnsCache = newDNSMap()
		return nil
	}

	return dnsCache.get(name, qtype)
}

func localResolver(name string) []string {
	addrs, err := net.LookupHost(name)
	if err != nil {
		return []string{}
	}

	return addrs
}

func mmesh64Resolver(name string) []string {
	s := strings.Split(name, ".")

	// resolver for mmesh64 names: 1.2.3.4.<gw>.mmesh.local.

	if len(s) != 8 {
		return []string{}
	}

	if s[5] == "mmesh" && s[6] == "local" {
		ipv4 := fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
		if net.ParseIP(ipv4) != nil {
			addr, err := ipnet.GetMMesh64Addr(mnet.LocalNode().Router().IPv6(), ipv4)
			if err != nil {
				return []string{}
			}

			return []string{addr}
		}
	}

	return []string{}
}

/*
func dnsProxy(q []dns.Question) (*dns.Msg, error) {
	if len(q) == 0 {
		return nil, errors.New("invalid DNS question")
	}

	m := new(dns.Msg)
	m.RecursionDesired = true
	m.Question = q

	return dns.Exchange(m, dnsForwarder)
}
*/
