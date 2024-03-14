package netflowdb

import (
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
)

func aggregateByProtocol(tbp []*netdb.TrafficByProtocol) []*netdb.TrafficByProtocol {
	if len(tbp) <= maxEntries {
		return tbp
	}

	tmStart := tbp[0].Timestamp
	tmEnd := tbp[len(tbp)-1].Timestamp

	aggInterval := (tmEnd - tmStart) / maxEntries

	tbpm := make(map[int64]*netdb.TrafficByProtocol, 0)

	for i := 0; i < maxEntries; i++ {
		tm := tmStart + (int64(i) * aggInterval)

		tbpm[tm] = &netdb.TrafficByProtocol{
			Timestamp: tm + (aggInterval / 2),
			Traffic:   make(map[string]*netdb.TrafficCounter, 0),
		}
	}

	for tm, aggTBP := range tbpm {
		for _, tbpe := range tbp {
			if tbpe.Timestamp < tm {
				continue
			}

			if tbpe.Timestamp > (tm + aggInterval) {
				continue
			}

			for proto, tc := range tbpe.Traffic {
				if _, ok := aggTBP.Traffic[proto]; !ok {
					aggTBP.Traffic[proto] = &netdb.TrafficCounter{}
				}

				aggTBP.Traffic[proto].Bytes += tc.Bytes
				aggTBP.Traffic[proto].Packets += tc.Packets
			}
		}
	}

	ntbp := make([]*netdb.TrafficByProtocol, 0)

	for _, aggTBP := range tbpm {
		isZero := true

		for _, tc := range aggTBP.Traffic {
			if tc.Bytes > 0 {
				isZero = false
				break
			}
		}

		if isZero {
			continue
		}

		ntbp = append(ntbp, aggTBP)
	}

	return ntbp
}

/*
func aggregateAllProtocols(tbp []*netdb.TrafficByProtocol) []*netdb.TrafficByProtocol {
	ntbp := aggregateByProtocol(tbp)

	sort.Slice(ntbp, func(i, j int) bool {
		return ntbp[i].Timestamp < ntbp[j].Timestamp
	})

	return ntbp
}
*/

func protocols() []netdb.Protocol {
	return []netdb.Protocol{
		netdb.Protocol_UNKNOWN_PROTO,
		netdb.Protocol_TCP,
		netdb.Protocol_UDP,
		netdb.Protocol_ICMP4,
		netdb.Protocol_ICMP6,
		netdb.Protocol_GRE,
		netdb.Protocol_SCTP,
	}
}
