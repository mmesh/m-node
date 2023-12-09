package conntrack

import "github.com/google/gopacket/layers"

// Filter ICMP Type EchoRequest
func (conn *Connection) invalidICMPTypeRequest() bool {
	switch conn.Proto {
	case layers.IPProtocolICMPv4:
		if conn.ICMPv4TypeCode.Type() != layers.ICMPv4TypeEchoRequest {
			return true // only icmp echo request is permitted, drop the pkt
		}
	case layers.IPProtocolICMPv6:
		if conn.ICMPv6TypeCode.Type() != layers.ICMPv6TypeEchoRequest {
			return true // only icmp echo request is permitted, drop the pkt
		}
	}

	return false // icmp type is valid, accept the pkt
}

// Filter ICMP Type EchoReply
func (conn *Connection) invalidICMPTypeReply() bool {
	switch conn.Proto {
	case layers.IPProtocolICMPv4:
		if conn.ICMPv4TypeCode.Type() != layers.ICMPv4TypeEchoReply {
			return true // only icmp echo reply is permitted, drop the pkt
		}
	case layers.IPProtocolICMPv6:
		if conn.ICMPv6TypeCode.Type() != layers.ICMPv6TypeEchoReply {
			return true // only icmp echo reply is permitted, drop the pkt
		}
	}

	return false // icmp type is valid, accept the pkt
}
