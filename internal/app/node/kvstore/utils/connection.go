package utils

import (
	"fmt"
	"strconv"

	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-lib/pkg/errors"
)

func ParseAddressFamily(str string) (netdb.AddressFamily, error) {
	if len(str) == 0 {
		return netdb.AddressFamily_UNKNOWN_AF, fmt.Errorf("[netdb] invalid address family")
	}

	af, err := strconv.Atoi(str)
	if err != nil {
		return netdb.AddressFamily_UNKNOWN_AF, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	switch af {
	case int(netdb.AddressFamily_IP4):
		return netdb.AddressFamily_IP4, nil
	case int(netdb.AddressFamily_IP6):
		return netdb.AddressFamily_IP6, nil
	}

	return netdb.AddressFamily_UNKNOWN_AF, fmt.Errorf("[netdb] unknown address family")
}

func ParseProto(str string) (netdb.Protocol, error) {
	if len(str) == 0 {
		return netdb.Protocol_UNKNOWN_PROTO, fmt.Errorf("[netdb] invalid protocol")
	}

	proto, err := strconv.Atoi(str)
	if err != nil {
		return netdb.Protocol_UNKNOWN_PROTO, errors.Wrapf(err, "[%v] function strconv.Atoi()", errors.Trace())
	}

	switch proto {
	case int(netdb.Protocol_TCP):
		return netdb.Protocol_TCP, nil
	case int(netdb.Protocol_UDP):
		return netdb.Protocol_UDP, nil
	case int(netdb.Protocol_ICMP4):
		return netdb.Protocol_ICMP4, nil
	case int(netdb.Protocol_ICMP6):
		return netdb.Protocol_ICMP6, nil
	case int(netdb.Protocol_GRE):
		return netdb.Protocol_GRE, nil
	case int(netdb.Protocol_SCTP):
		return netdb.Protocol_SCTP, nil
	}

	return netdb.Protocol_UNKNOWN_PROTO, fmt.Errorf("[netdb] unknown protocol")
}
