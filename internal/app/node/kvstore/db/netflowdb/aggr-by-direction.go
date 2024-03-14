package netflowdb

import (
	"mmesh.dev/m-api-go/grpc/resources/nstore/netdb"
)

func aggregateByDirection(tbd []*netdb.TrafficByDirection) []*netdb.TrafficByDirection {
	if len(tbd) <= maxEntries {
		return tbd
	}

	tmStart := tbd[0].Timestamp
	tmEnd := tbd[len(tbd)-1].Timestamp

	aggInterval := (tmEnd - tmStart) / maxEntries

	tbdm := make(map[int64]*netdb.TrafficByDirection, 0)

	for i := 0; i < maxEntries; i++ {
		tm := tmStart + (int64(i) * aggInterval)

		tbdm[tm] = &netdb.TrafficByDirection{
			Timestamp: tm + (aggInterval / 2),
			Traffic:   make(map[string]*netdb.TrafficCounter, 0),
		}
	}

	for tm, aggTBD := range tbdm {
		for _, tbde := range tbd {
			if tbde.Timestamp < tm {
				continue
			}

			if tbde.Timestamp > (tm + aggInterval) {
				continue
			}

			for dir, tc := range tbde.Traffic {
				if _, ok := aggTBD.Traffic[dir]; !ok {
					aggTBD.Traffic[dir] = &netdb.TrafficCounter{}
				}

				aggTBD.Traffic[dir].Bytes += tc.Bytes
				aggTBD.Traffic[dir].Packets += tc.Packets
			}
		}
	}

	ntbd := make([]*netdb.TrafficByDirection, 0)

	for _, aggTBD := range tbdm {
		isZero := true

		for _, tc := range aggTBD.Traffic {
			if tc.Bytes > 0 {
				isZero = false
				break
			}
		}

		if isZero {
			continue
		}

		ntbd = append(ntbd, aggTBD)
	}

	return ntbd
}

/*
func aggregateByDirection(dir netdb.ConnectionDirection, tbd []*netdb.TrafficByDirection) []*netdb.TrafficByDirection {
	if len(tbd) <= maxEntries {
		return tbd
	}

	tmStart := tbd[0].Timestamp
	tmEnd := tbd[len(tbd)-1].Timestamp

	aggInterval := (tmEnd - tmStart) / maxEntries

	tbdm := make(map[int64]*netdb.TrafficByDirection, 0)

	for i := 0; i < maxEntries; i++ {
		tm := tmStart + (int64(i) * aggInterval)

		tbdm[tm] = &netdb.TrafficByDirection{
			Timestamp: tm + (aggInterval / 2),
			Traffic:   &netdb.TrafficCounter{},
			Direction: dir,
		}
	}

	for tm, aggTBD := range tbdm {
		for _, tbde := range tbd {
			if tbde.Direction != dir {
				continue
			}

			if tbde.Timestamp < tm {
				continue
			}

			if tbde.Timestamp > (tm + aggInterval) {
				continue
			}

			aggTBD.Traffic.Bytes += tbde.Traffic.Bytes
			aggTBD.Traffic.Packets += tbde.Traffic.Packets
		}
	}

	ntbd := make([]*netdb.TrafficByDirection, 0)

	for _, aggTBD := range tbdm {
		if aggTBD.Traffic.Bytes == 0 {
			continue
		}

		ntbd = append(ntbd, aggTBD)
	}

	return ntbd
}

func aggregateAllDirections(tbd []*netdb.TrafficByDirection) []*netdb.TrafficByDirection {
	ntbd := make([]*netdb.TrafficByDirection, 0)

	for _, dir := range directions() {
		aggtbd := aggregateByDirection(dir, tbd)
		ntbd = append(ntbd, aggtbd...)
	}

	sort.Slice(ntbd, func(i, j int) bool {
		return ntbd[i].Timestamp < ntbd[j].Timestamp
	})

	return ntbd
}
*/

func directions() []netdb.ConnectionDirection {
	return []netdb.ConnectionDirection{
		netdb.ConnectionDirection_UNKNOWN_DIRECTION,
		netdb.ConnectionDirection_INCOMING,
		netdb.ConnectionDirection_OUTGOING,
	}
}
