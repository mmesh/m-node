package ipnet

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/viper"
)

const (
	fqdnIAPDomain    string = "iap.mmesh.io"
	fqdnIAPDomainDev string = "iap.dev.mmesh.io"
)

func IAPDomain() string {
	if viper.GetString("version.branch") == "dev" {
		return fqdnIAPDomainDev
	}

	return fqdnIAPDomain
}

func VSCNAMEIsValid(fqdn, locationID string) error {
	if len(fqdn) == 0 {
		return fmt.Errorf("missing fqdn")
	}

	if len(locationID) == 0 {
		return fmt.Errorf("missing locationID")
	}

	vHost := fmt.Sprintf("vs.%s.%s", locationID, IAPDomain())

	if fqdn == vHost {
		return fmt.Errorf("invalid cname")
	}

	cname, err := net.LookupCNAME(fqdn)
	if err != nil {
		return fmt.Errorf("invalid cname: %v", err)
	}

	cname = strings.TrimSuffix(cname, ".")

	if cname != vHost {
		return fmt.Errorf("CNAME does not match target: (%s != %s)", cname, vHost)
	}

	return nil
}
