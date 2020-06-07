package k8s

import (
	"github.com/spf13/viper"
	v1 "k8s.io/api/core/v1"
)

type k8sSvcAnnotationsCfg struct {
	//enabled bool
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

	cfgAccountID := viper.GetString("node.account")
	cfgTenantID := viper.GetString("node.tenant")
	cfgNetID := viper.GetString("node.network")
	cfgVRFID := viper.GetString("node.subnet")

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

	vrfID, ok := s.ObjectMeta.Annotations["mmesh.io/subnet"]
	if !ok {
		valid = false
	}
	if vrfID != cfgVRFID {
		valid = false
	}

	dnsName, ok := s.ObjectMeta.Annotations["mmesh.io/dnsName"]
	if !ok {
		dnsName = s.ObjectMeta.Name
	}

	reqIPv4, ok := s.ObjectMeta.Annotations["mmesh.io/ipv4"]
	if !ok {
		reqIPv4 = "auto"
	}

	return &k8sSvcAnnotationsCfg{
		//enabled: svcEnabled,
		dnsName: dnsName,
		reqIPv4: reqIPv4,
		valid:   valid,
	}
}
