package netflowdb

import (
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
)

func aggregateByAddressFamily(tbaf []*netdb.TrafficByAddressFamily) []*netdb.TrafficByAddressFamily {
	if len(tbaf) <= maxEntries {
		return tbaf
	}

	tmStart := tbaf[0].Timestamp
	tmEnd := tbaf[len(tbaf)-1].Timestamp

	aggInterval := (tmEnd - tmStart) / maxEntries

	tbafm := make(map[int64]*netdb.TrafficByAddressFamily, 0)

	for i := 0; i < maxEntries; i++ {
		tm := tmStart + (int64(i) * aggInterval)

		tbafm[tm] = &netdb.TrafficByAddressFamily{
			Timestamp: tm + (aggInterval / 2),
			Traffic:   make(map[string]*netdb.TrafficCounter, 0),
		}
	}

	for tm, aggTBAF := range tbafm {
		for _, tbafe := range tbaf {
			if tbafe.Timestamp < tm {
				continue
			}

			if tbafe.Timestamp > (tm + aggInterval) {
				continue
			}

			for dir, tc := range tbafe.Traffic {
				if _, ok := aggTBAF.Traffic[dir]; !ok {
					aggTBAF.Traffic[dir] = &netdb.TrafficCounter{}
				}

				aggTBAF.Traffic[dir].Bytes += tc.Bytes
				aggTBAF.Traffic[dir].Packets += tc.Packets
			}
		}
	}

	ntbaf := make([]*netdb.TrafficByAddressFamily, 0)

	for _, aggTBAF := range tbafm {
		isZero := true

		for _, tc := range aggTBAF.Traffic {
			if tc.Bytes > 0 {
				isZero = false
				break
			}
		}

		if isZero {
			continue
		}

		ntbaf = append(ntbaf, aggTBAF)
	}

	return ntbaf
}

func addressFamilies() []netdb.AddressFamily {
	return []netdb.AddressFamily{
		netdb.AddressFamily_UNKNOWN_AF,
		netdb.AddressFamily_IP4,
		netdb.AddressFamily_IP6,
	}
}
