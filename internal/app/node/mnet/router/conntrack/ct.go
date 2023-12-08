package conntrack

import (
	"time"
)

type ctStatus uint32

const (
	ctStatusExpected ctStatus = iota
	ctStatusSeenReply
	// ctStatusAsured
	// ctStatusConfirmed
)

const ctTimeout = 120

type ctState struct {
	status        ctStatus
	timeout       time.Time
	originCounter *ctCounter
	replyCounter  *ctCounter
	// protoInfo     *protoInfo
}

type ctCounter struct {
	packets uint64
	bytes   uint64
}

/*
type protoInfo struct {
	tcp  *tcpInfo
	dccp *dccpInfo
	sctp *sctpInfo
	gre  *greInfo
}

type tcpInfo struct {
	state      *uint8
	wScaleOrig *uint8
	wScaleRepl *uint8
	flagsOrig  *tcpFlags
	flagsReply *tcpFlags
}

type tcpFlags struct {
	flags *uint8
	mask  *uint8
}

type dccpInfo struct {
	state        *uint8
	role         *uint8
	handshakeSeq *uint64
}

type sctpInfo struct {
	state        *uint8
	vTagOriginal *uint32
	vTagReply    *uint32
}

type greInfo struct {
	key *uint32
}
*/

/*
type timestamp struct {
	start *time.Time
	stop  *time.Time
}

type ipTuple struct {
	srcIP *netip.Addr
	dstIP *netip.Addr
	proto *ipProto
	// Zone  uint16
}

type ipProto struct {
	proto         layers.IPProtocol
	srcPort       uint16
	dstPort       uint16
	icmp4ID       uint16
	icmp4TypeCode layers.ICMPv4TypeCode
	icmp6ID       uint16
	icmp6TypeCode layers.ICMPv6TypeCode
}

type ctFlow struct {
	id        uint32
	status    connStatus
	timeout   uint32
	origin    *ipTuple
	reply     *ipTuple
	protoInfo *protoInfo
}
*/
