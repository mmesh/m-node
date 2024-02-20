package ipnet

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/mdlayher/netx/eui64"
	"mmesh.dev/m-lib/pkg/errors"
)

type AddressFamily int32

const (
	AddressFamilyUnspec AddressFamily = 0
	AddressFamilyIPv4   AddressFamily = 4
	AddressFamilyIPv6   AddressFamily = 6
)

const mmeshIPv6Prefix string = "fd77:f:"

func (af AddressFamily) String() string {
	switch af {
	case AddressFamilyIPv4:
		return "IPv4"
	case AddressFamilyIPv6:
		return "IPv6"
	}

	return "unspec"
}

func IPv6Prefix() string {
	return mmeshIPv6Prefix
}

func GetIPv6(ipv4 string) (string, error) {
	ip := strings.Split(ipv4, ".")

	if len(ip) != 4 {
		return "", errors.Errorf("invalid ipv4: %s", ipv4)
	}

	a, err := strconv.Atoi(ip[0])
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	b, err := strconv.Atoi(ip[1])
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	c, err := strconv.Atoi(ip[2])
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	d, err := strconv.Atoi(ip[3])
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	ipv6 := fmt.Sprintf("%s%02x%02x:%02x%02x:1:0:0:0", mmeshIPv6Prefix, a, b, c, d)

	return net.ParseIP(ipv6).String(), nil
}

func GetIPv6Endpoint(ipv6 string) (string, error) {
	if !strings.HasPrefix(ipv6, mmeshIPv6Prefix) {
		return "", errors.Errorf("invalid mmesh64 ipv6: %s", ipv6)
	}

	s := strings.Split(ipv6, ":")

	if len(s) < 6 {
		return "", errors.Errorf("invalid mmesh64 ipv6: %s", ipv6)
	}

	addr := fmt.Sprintf("%s%s:%s:1:0:0:0", mmeshIPv6Prefix, s[2], s[3])

	return net.ParseIP(addr).String(), nil
}

func GetMMesh64Addr(ipv6, ipv4 string) (string, error) {
	if !strings.HasPrefix(ipv6, mmeshIPv6Prefix) {
		return "", errors.Errorf("invalid mmesh64 ipv6: %s", ipv6)
	}

	s1 := strings.Split(ipv6, ":")
	s2 := strings.Split(ipv4, ".")

	if len(s1) < 6 {
		return "", errors.Errorf("invalid mmesh64 ipv6: %s", ipv6)
	}

	if len(s2) != 4 {
		return "", errors.Errorf("invalid ipv4: %s", ipv4)
	}

	addr := fmt.Sprintf("%s%s:%s:%s:%s:%s:%s", mmeshIPv6Prefix, s1[2], s1[3], s2[0], s2[1], s2[2], s2[3])

	return net.ParseIP(addr).String(), nil
}

func GetIPv4Encap(ipv6 string) (string, error) {
	if !strings.HasPrefix(ipv6, mmeshIPv6Prefix) {
		return "", errors.Errorf("invalid mmesh64 ipv6: %s", ipv6)
	}

	s := strings.Split(ipv6, ":")

	if len(s) != 8 {
		return "", errors.Errorf("this ipv6 addr does not encapsulate an ipv4 addr: %s", ipv6)
	}

	ipv4 := fmt.Sprintf("%s.%s.%s.%s", s[4], s[5], s[6], s[7])

	return net.ParseIP(ipv4).String(), nil
}

func IsIAPIPv6Addr(ipv6 string) bool {
	prefix1 := fmt.Sprintf("%s:", mmeshIPv6Prefix)
	prefix2 := fmt.Sprintf("%s0:0:0:0:", mmeshIPv6Prefix)

	return strings.HasPrefix(ipv6, prefix1) || strings.HasPrefix(ipv6, prefix2)
}

func GetIPv6ByMac(ipv6 string) (net.IP, error) {
	if !strings.HasPrefix(ipv6, mmeshIPv6Prefix) {
		return nil, errors.Errorf("invalid mmesh64 ipv6: %s", ipv6)
	}

	s := strings.Split(ipv6, ":")

	if len(s) < 6 {
		return nil, errors.Errorf("invalid mmesh64 ipv6: %s", ipv6)
	}

	prefix := fmt.Sprintf("%s%s:%s::", mmeshIPv6Prefix, s[2], s[3])

	mmeshPrefix := net.ParseIP(prefix)

	mac, err := getInterfaceHwAddr()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getInterfaceHwAddr()", errors.Trace())
	}

	// Retrieve IPv6 address from IPv6 prefix and MAC address
	addrv6, err := eui64.ParseMAC(mmeshPrefix, mac)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function eui64.ParseMAC()", errors.Trace())
	}

	return addrv6, nil
}
