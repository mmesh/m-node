package k8s

import (
	v1 "k8s.io/api/core/v1"
	"mmesh.dev/m-node/internal/app/node/mnet"
	"mmesh.dev/m-node/internal/app/node/utils"
)

type k8sSvcAnnotationsCfg struct {
	// enabled bool
	dnsName string
	reqIPv4 string
	valid   bool
}

func parseAnnotations(s *v1.Service) *k8sSvcAnnotationsCfg {
	// var svcEnabled bool
	// var err error

	// if enabled, ok := s.ObjectMeta.Annotations["mmesh.io/enabled"]; !ok {
	// 	return nil
	// } else {
	// 	if svcEnabled, err = strconv.ParseBool(enabled); err != nil {
	// 		return nil
	// 	}
	// }

	n := mnet.LocalNode().Node()

	cfgAccountID := n.AccountID
	cfgTenantID := n.TenantID
	cfgNetID := n.Cfg.NetID
	cfgSubnetID := n.Cfg.SubnetID

	valid := true

	accountID, ok := s.ObjectMeta.Annotations["mmesh.io/account"]
	if !ok {
		valid = false
	}
	if accountID != cfgAccountID {
		valid = false
	}

	tenantID, ok := s.ObjectMeta.Annotations["mmesh.io/tenant"]
	if !ok {
		valid = false
	}
	if tenantID != cfgTenantID {
		valid = false
	}

	netID, ok := s.ObjectMeta.Annotations["mmesh.io/network"]
	if !ok {
		valid = false
	}
	if netID != cfgNetID {
		valid = false
	}

	subnetID, ok := s.ObjectMeta.Annotations["mmesh.io/subnet"]
	if !ok {
		valid = false
	}
	if subnetID != cfgSubnetID {
		valid = false
	}

	dnsName, ok := s.ObjectMeta.Annotations["mmesh.io/dnsName"]
	if !ok {
		dnsName = s.ObjectMeta.Name
	}

	reqIPv4, ok := s.ObjectMeta.Annotations["mmesh.io/ipv4"]
	if ok {
		if !utils.IPv4IsValid(reqIPv4) {
			reqIPv4 = "auto"
		}
	} else {
		reqIPv4 = "auto"
	}

	return &k8sSvcAnnotationsCfg{
		// enabled: svcEnabled,
		dnsName: dnsName,
		reqIPv4: reqIPv4,
		valid:   valid,
	}
}
