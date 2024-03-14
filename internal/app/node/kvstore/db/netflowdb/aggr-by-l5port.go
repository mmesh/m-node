package netflowdb

import (
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
)

func aggregateByL5Port(tbl5p []*netdb.TrafficByL5Port) []*netdb.TrafficByL5Port {
	if len(tbl5p) <= maxEntries {
		return tbl5p
	}

	tmStart := tbl5p[0].Timestamp
	tmEnd := tbl5p[len(tbl5p)-1].Timestamp

	aggInterval := (tmEnd - tmStart) / maxEntries

	tbl5pm := make(map[int64]*netdb.TrafficByL5Port, 0)

	for i := 0; i < maxEntries; i++ {
		tm := tmStart + (int64(i) * aggInterval)

		tbl5pm[tm] = &netdb.TrafficByL5Port{
			Timestamp: tm + (aggInterval / 2),
			Traffic:   make(map[string]*netdb.TrafficCounter, 0),
		}
	}

	for tm, aggTBL5P := range tbl5pm {
		for _, tbl5pe := range tbl5p {
			if tbl5pe.Timestamp < tm {
				continue
			}

			if tbl5pe.Timestamp > (tm + aggInterval) {
				continue
			}

			for l5port, tc := range tbl5pe.Traffic {
				if _, ok := aggTBL5P.Traffic[l5port]; !ok {
					aggTBL5P.Traffic[l5port] = &netdb.TrafficCounter{}
				}

				aggTBL5P.Traffic[l5port].Bytes += tc.Bytes
				aggTBL5P.Traffic[l5port].Packets += tc.Packets
			}
		}
	}

	ntbl5p := make([]*netdb.TrafficByL5Port, 0)

	for _, aggTBL5P := range tbl5pm {
		isZero := true

		for _, tc := range aggTBL5P.Traffic {
			if tc.Bytes > 0 {
				isZero = false
				break
			}
		}

		if isZero {
			continue
		}

		ntbl5p = append(ntbl5p, aggTBL5P)
	}

	return ntbl5p
}

/*
func aggregateByL5Port(l5port netdb.L5Port, tbl5p []*netdb.TrafficByL5Port) []*netdb.TrafficByL5Port {
	if len(tbl5p) <= maxEntries {
		return tbl5p
	}

	tmStart := tbl5p[0].Timestamp
	tmEnd := tbl5p[len(tbl5p)-1].Timestamp

	aggInterval := (tmEnd - tmStart) / maxEntries

	tbl5pm := make(map[int64]*netdb.TrafficByL5Port, 0)

	for i := 0; i < maxEntries; i++ {
		tm := tmStart + (int64(i) * aggInterval)

		tbl5pm[tm] = &netdb.TrafficByL5Port{
			Timestamp: tm + (aggInterval / 2),
			Traffic:   &netdb.TrafficCounter{},
			L5Port:    l5port,
		}
	}

	for tm, aggTBL5P := range tbl5pm {
		for _, tbl5pe := range tbl5p {
			if tbl5pe.L5Port != l5port {
				continue
			}

			if tbl5pe.Timestamp < tm {
				continue
			}

			if tbl5pe.Timestamp > (tm + aggInterval) {
				continue
			}

			aggTBL5P.Traffic.Bytes += tbl5pe.Traffic.Bytes
			aggTBL5P.Traffic.Packets += tbl5pe.Traffic.Packets
		}
	}

	ntbl5p := make([]*netdb.TrafficByL5Port, 0)

	for _, aggTBL5P := range tbl5pm {
		if aggTBL5P.Traffic.Bytes == 0 {
			continue
		}

		ntbl5p = append(ntbl5p, aggTBL5P)
	}

	return ntbl5p
}

func aggregateAllL5Ports(tbl5p []*netdb.TrafficByL5Port) []*netdb.TrafficByL5Port {
	ntbl5p := make([]*netdb.TrafficByL5Port, 0)

	for _, port := range l5ports() {
		aggtbl5p := aggregateByL5Port(port, tbl5p)
		ntbl5p = append(ntbl5p, aggtbl5p...)
	}

	sort.Slice(ntbl5p, func(i, j int) bool {
		return ntbl5p[i].Timestamp < ntbl5p[j].Timestamp
	})

	return ntbl5p
}
*/

func l5ports() []netdb.L5Port {
	return []netdb.L5Port{
		netdb.L5Port_OTHER_L5PORT,
		netdb.L5Port_HTTP,
		netdb.L5Port_HTTPS,
		netdb.L5Port_SSH,
		netdb.L5Port_RDP,
		netdb.L5Port_SMB,
	}
}
