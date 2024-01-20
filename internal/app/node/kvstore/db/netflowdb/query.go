package netflowdb

import (
	"sort"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
	"mmesh.dev/m-lib/pkg/errors"
)

func newTimeTrafficValue(tm int64, traffic *netdb.TrafficCounter) *netdb.TimeTrafficValue {
	return &netdb.TimeTrafficValue{
		Timestamp: tm,
		Traffic:   traffic,
	}
}

func (nfdb *netflowDB) Query(r *netdb.TrafficMetricsRequest) (*netdb.TrafficMetricsResponse, error) {
	tmr := &netdb.TrafficMetricsResponse{
		AccountID:   r.Request.AccountID,
		TenantID:    r.Request.TenantID,
		NodeID:      r.Request.NodeID,
		QueryID:     r.Request.QueryID,
		ByProtocol:  nil,
		ByL5Port:    nil,
		ByDirection: nil,
		TopTalkers:  nil,
		Timestamp:   time.Now().UnixMilli(),
	}

	nfl, err := nfdb.Scan()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function nfdb.Scan()", errors.Trace())
	}

	switch r.Type {
	case netdb.TrafficQueryType_TRAFFIC_BY_PROTOCOL:
		tmr.ByProtocol = getTrafficByProtocol(nfl)
	case netdb.TrafficQueryType_TRAFFIC_BY_L5_PORT:
		tmr.ByL5Port = getTrafficByL5Port(nfl)
	case netdb.TrafficQueryType_TRAFFIC_BY_DIRECTION:
		tmr.ByDirection = getTrafficByDirection(nfl)
	case netdb.TrafficQueryType_TRAFFIC_TOP_TALKERS:
		tmr.TopTalkers = getTrafficTopTalkers(nfl)
	}

	return tmr, nil
}

func getTrafficByProtocol(nfl []*netdb.NetFlowEntry) *netdb.TrafficByProtocol {
	tbp := &netdb.TrafficByProtocol{
		Unknown: make([]*netdb.TimeTrafficValue, 0),
		TCP:     make([]*netdb.TimeTrafficValue, 0),
		UDP:     make([]*netdb.TimeTrafficValue, 0),
		ICMPv4:  make([]*netdb.TimeTrafficValue, 0),
		ICMPv6:  make([]*netdb.TimeTrafficValue, 0),
		GRE:     make([]*netdb.TimeTrafficValue, 0),
		SCTP:    make([]*netdb.TimeTrafficValue, 0),
	}

	for _, nfe := range nfl {
		if nfe.Flow == nil || nfe.Flow.Connection == nil {
			continue
		}

		switch nfe.Flow.Connection.Proto {
		case netdb.Protocol_UNKNOWN_PROTO:
			tbp.Unknown = append(tbp.Unknown, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		case netdb.Protocol_TCP:
			tbp.TCP = append(tbp.TCP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		case netdb.Protocol_UDP:
			tbp.UDP = append(tbp.UDP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		case netdb.Protocol_ICMP4:
			tbp.ICMPv4 = append(tbp.ICMPv4, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		case netdb.Protocol_ICMP6:
			tbp.ICMPv6 = append(tbp.ICMPv6, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		case netdb.Protocol_GRE:
			tbp.GRE = append(tbp.GRE, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		case netdb.Protocol_SCTP:
			tbp.SCTP = append(tbp.SCTP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		}
	}

	return tbp
}

func getTrafficByL5Port(nfl []*netdb.NetFlowEntry) *netdb.TrafficByL5Port {
	tbl5p := &netdb.TrafficByL5Port{
		Other:          make([]*netdb.TimeTrafficValue, 0),
		HTTP:           make([]*netdb.TimeTrafficValue, 0),
		HTTPS:          make([]*netdb.TimeTrafficValue, 0),
		SSH:            make([]*netdb.TimeTrafficValue, 0),
		RDP:            make([]*netdb.TimeTrafficValue, 0),
		DNS:            make([]*netdb.TimeTrafficValue, 0),
		SMTP:           make([]*netdb.TimeTrafficValue, 0),
		SMTPS:          make([]*netdb.TimeTrafficValue, 0),
		MailSubmission: make([]*netdb.TimeTrafficValue, 0),
		IMAP:           make([]*netdb.TimeTrafficValue, 0),
		IMAPS:          make([]*netdb.TimeTrafficValue, 0),
		POP3:           make([]*netdb.TimeTrafficValue, 0),
		POP3S:          make([]*netdb.TimeTrafficValue, 0),
		NTP:            make([]*netdb.TimeTrafficValue, 0),
		SNMP:           make([]*netdb.TimeTrafficValue, 0),
		BGP:            make([]*netdb.TimeTrafficValue, 0),
		LDAP:           make([]*netdb.TimeTrafficValue, 0),
		LDAPS:          make([]*netdb.TimeTrafficValue, 0),
		MySQL:          make([]*netdb.TimeTrafficValue, 0),
		PostgreSQL:     make([]*netdb.TimeTrafficValue, 0),
		MSSQL:          make([]*netdb.TimeTrafficValue, 0),
		Redis:          make([]*netdb.TimeTrafficValue, 0),
		NFS:            make([]*netdb.TimeTrafficValue, 0),
		SIP:            make([]*netdb.TimeTrafficValue, 0),
		SIPTLS:         make([]*netdb.TimeTrafficValue, 0),
		AMQP:           make([]*netdb.TimeTrafficValue, 0),
		AMQPS:          make([]*netdb.TimeTrafficValue, 0),
	}

	for _, nfe := range nfl {
		if nfe.Flow == nil || nfe.Flow.Connection == nil {
			continue
		}

		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 80 || nfe.Flow.Connection.SrcPort == 80) {
			tbl5p.HTTP = append(tbl5p.HTTP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 8080 || nfe.Flow.Connection.SrcPort == 8080) {
			tbl5p.HTTP = append(tbl5p.HTTP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 443 || nfe.Flow.Connection.SrcPort == 443 {
			tbl5p.HTTPS = append(tbl5p.HTTPS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 22 || nfe.Flow.Connection.SrcPort == 22) {
			tbl5p.SSH = append(tbl5p.SSH, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 3389 || nfe.Flow.Connection.SrcPort == 3389) {
			tbl5p.RDP = append(tbl5p.RDP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 53 || nfe.Flow.Connection.SrcPort == 53 {
			tbl5p.DNS = append(tbl5p.DNS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 25 || nfe.Flow.Connection.SrcPort == 25) {
			tbl5p.SMTP = append(tbl5p.SMTP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 465 || nfe.Flow.Connection.SrcPort == 465) {
			tbl5p.SMTPS = append(tbl5p.SMTPS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 587 || nfe.Flow.Connection.SrcPort == 587) {
			tbl5p.MailSubmission = append(tbl5p.MailSubmission, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 143 || nfe.Flow.Connection.SrcPort == 143) {
			tbl5p.IMAP = append(tbl5p.IMAP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 993 || nfe.Flow.Connection.SrcPort == 993) {
			tbl5p.IMAPS = append(tbl5p.IMAPS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 110 || nfe.Flow.Connection.SrcPort == 110) {
			tbl5p.POP3 = append(tbl5p.POP3, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 995 || nfe.Flow.Connection.SrcPort == 995) {
			tbl5p.POP3S = append(tbl5p.POP3S, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_UDP &&
			(nfe.Flow.Connection.DstPort == 123 || nfe.Flow.Connection.SrcPort == 123) {
			tbl5p.NTP = append(tbl5p.NTP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 161 || nfe.Flow.Connection.SrcPort == 161 {
			tbl5p.SNMP = append(tbl5p.SNMP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 179 || nfe.Flow.Connection.SrcPort == 179) {
			tbl5p.BGP = append(tbl5p.BGP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 389 || nfe.Flow.Connection.SrcPort == 389 {
			tbl5p.LDAP = append(tbl5p.LDAP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 636 || nfe.Flow.Connection.SrcPort == 636 {
			tbl5p.LDAPS = append(tbl5p.LDAPS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 3306 || nfe.Flow.Connection.SrcPort == 3306) {
			tbl5p.MySQL = append(tbl5p.MySQL, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 5432 || nfe.Flow.Connection.SrcPort == 5432) {
			tbl5p.PostgreSQL = append(tbl5p.PostgreSQL, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 1433 || nfe.Flow.Connection.SrcPort == 1433) {
			tbl5p.MSSQL = append(tbl5p.MSSQL, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 6379 || nfe.Flow.Connection.SrcPort == 6379) {
			tbl5p.Redis = append(tbl5p.Redis, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 2049 || nfe.Flow.Connection.SrcPort == 2049 {
			tbl5p.NFS = append(tbl5p.NFS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 5060 || nfe.Flow.Connection.SrcPort == 5060 {
			tbl5p.SIP = append(tbl5p.SIP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.DstPort == 5061 || nfe.Flow.Connection.SrcPort == 5061 {
			tbl5p.SIPTLS = append(tbl5p.SIPTLS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 5672 || nfe.Flow.Connection.SrcPort == 5672) {
			tbl5p.AMQP = append(tbl5p.AMQP, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}
		if nfe.Flow.Connection.Proto == netdb.Protocol_TCP &&
			(nfe.Flow.Connection.DstPort == 5671 || nfe.Flow.Connection.SrcPort == 5671) {
			tbl5p.AMQPS = append(tbl5p.AMQPS, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
			continue
		}

		tbl5p.Other = append(tbl5p.Other, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
	}

	return tbl5p
}

func getTrafficByDirection(nfl []*netdb.NetFlowEntry) *netdb.TrafficByDirection {
	tbd := &netdb.TrafficByDirection{
		Incoming: make([]*netdb.TimeTrafficValue, 0),
		Outgoing: make([]*netdb.TimeTrafficValue, 0),
	}

	for _, nfe := range nfl {
		if nfe.Flow == nil || nfe.Flow.Connection == nil {
			continue
		}

		switch nfe.Flow.Direction {
		case netdb.ConnectionDirection_INCOMING:
			tbd.Incoming = append(tbd.Incoming, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		case netdb.ConnectionDirection_OUTGOING:
			tbd.Outgoing = append(tbd.Outgoing, newTimeTrafficValue(nfe.Timestamp, nfe.Traffic))
		}
	}

	return tbd
}

func getTrafficTopTalkers(nfl []*netdb.NetFlowEntry) *netdb.TopTalkers {
	srcTalkersMap := make(map[string]uint64, 0) // map[addr]bytes
	dstTalkersMap := make(map[string]uint64, 0) // map[addr]bytes

	for _, nfe := range nfl {
		if nfe.Flow == nil || nfe.Flow.Connection == nil {
			continue
		}

		switch nfe.Flow.Direction {
		case netdb.ConnectionDirection_INCOMING:
			srcTalkersMap[nfe.Flow.Connection.SrcIP] += nfe.Traffic.Bytes
		case netdb.ConnectionDirection_OUTGOING:
			dstTalkersMap[nfe.Flow.Connection.DstIP] += nfe.Traffic.Bytes
		}
	}

	return &netdb.TopTalkers{
		Src: getTopTalkers(srcTalkersMap),
		Dst: getTopTalkers(dstTalkersMap),
	}
}

func getTopTalkers(talkersMap map[string]uint64) []*netdb.Talker {
	talkers := make([]*netdb.Talker, 0)

	for addr, bytes := range talkersMap {
		talkers = append(talkers, &netdb.Talker{
			Addr:  addr,
			Bytes: bytes,
		})
	}

	sort.Slice(talkers, func(i, j int) bool {
		return talkers[i].Bytes > talkers[j].Bytes
	})

	n := 10
	if len(talkers) < n {
		n = len(talkers)
	}

	topTalkers := make([]*netdb.Talker, n)

	for i := 0; i < n; i++ {
		topTalkers[i] = talkers[i]
	}

	return topTalkers
}
