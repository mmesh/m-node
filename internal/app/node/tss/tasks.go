package tss

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/tss"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

type aggTask struct {
	srcTimeRange tss.TimeRange
	aggInterval  time.Duration
	jobInterval  time.Duration
}

var aggTaskMap map[tss.TimeRange]*aggTask = map[tss.TimeRange]*aggTask{
	tss.TimeRange_TTL_6H: {
		srcTimeRange: tss.TimeRange_TTL_1H,
		aggInterval:  12 * time.Minute,
		jobInterval:  12*time.Minute + (10 * time.Second),
	},
	tss.TimeRange_TTL_12H: {
		srcTimeRange: tss.TimeRange_TTL_1H,
		aggInterval:  24 * time.Minute,
		jobInterval:  24*time.Minute + (22 * time.Second),
	},
	tss.TimeRange_TTL_24H: {
		srcTimeRange: tss.TimeRange_TTL_6H,
		aggInterval:  48 * time.Minute,
		jobInterval:  48*time.Minute + (33 * time.Second),
	},
	tss.TimeRange_TTL_7D: {
		srcTimeRange: tss.TimeRange_TTL_12H,
		aggInterval:  336 * time.Minute,
		jobInterval:  336*time.Minute + (45 * time.Second),
	},
	tss.TimeRange_TTL_14D: {
		srcTimeRange: tss.TimeRange_TTL_24H,
		aggInterval:  672 * time.Minute,
		jobInterval:  672*time.Minute + (57 * time.Second),
	},
	tss.TimeRange_TTL_30D: {
		srcTimeRange: tss.TimeRange_TTL_7D,
		aggInterval:  1344 * time.Minute,
		jobInterval:  1344*time.Minute + (71 * time.Second),
	},
	tss.TimeRange_TTL_365D: {
		srcTimeRange: tss.TimeRange_TTL_30D,
		aggInterval:  12 * 24 * time.Hour,
		jobInterval:  12*24*time.Hour + (87 * time.Second),
	},
}

func (tsdb *tsDB) aggController() {
	ticker6H := time.NewTicker(aggTaskMap[tss.TimeRange_TTL_6H].jobInterval)
	defer ticker6H.Stop()
	ticker12H := time.NewTicker(aggTaskMap[tss.TimeRange_TTL_12H].jobInterval)
	defer ticker12H.Stop()
	ticker24H := time.NewTicker(aggTaskMap[tss.TimeRange_TTL_24H].jobInterval)
	defer ticker24H.Stop()
	ticker7D := time.NewTicker(aggTaskMap[tss.TimeRange_TTL_7D].jobInterval)
	defer ticker7D.Stop()
	ticker14D := time.NewTicker(aggTaskMap[tss.TimeRange_TTL_14D].jobInterval)
	defer ticker14D.Stop()
	ticker30D := time.NewTicker(aggTaskMap[tss.TimeRange_TTL_30D].jobInterval)
	defer ticker30D.Stop()
	ticker365D := time.NewTicker(aggTaskMap[tss.TimeRange_TTL_365D].jobInterval)
	defer ticker365D.Stop()

	for {
		select {
		case <-ticker6H.C:
			dstTimeRange := tss.TimeRange_TTL_6H
			if err := tsdb.runAggregation(dstTimeRange); err != nil {
				xlog.Errorf("[tss] Unable to complete aggregation job %s: %v",
					dstTimeRange.String(), err)
			}
			xlog.Debugf("[tss] Aggregation job %s completed", dstTimeRange.String())

		case <-ticker12H.C:
			dstTimeRange := tss.TimeRange_TTL_12H
			if err := tsdb.runAggregation(dstTimeRange); err != nil {
				xlog.Errorf("[tss] Unable to complete aggregation job %s: %v",
					dstTimeRange.String(), err)
			}
			xlog.Debugf("[tss] Aggregation job %s completed", dstTimeRange.String())

		case <-ticker24H.C:
			dstTimeRange := tss.TimeRange_TTL_24H
			if err := tsdb.runAggregation(dstTimeRange); err != nil {
				xlog.Errorf("[tss] Unable to complete aggregation job %s: %v",
					dstTimeRange.String(), err)
			}
			xlog.Debugf("[tss] Aggregation job %s completed", dstTimeRange.String())

		case <-ticker7D.C:
			dstTimeRange := tss.TimeRange_TTL_7D
			if err := tsdb.runAggregation(dstTimeRange); err != nil {
				xlog.Errorf("[tss] Unable to complete aggregation job %s: %v",
					dstTimeRange.String(), err)
			}
			xlog.Debugf("[tss] Aggregation job %s completed", dstTimeRange.String())

		case <-ticker14D.C:
			dstTimeRange := tss.TimeRange_TTL_14D
			if err := tsdb.runAggregation(dstTimeRange); err != nil {
				xlog.Errorf("[tss] Unable to complete aggregation job %s: %v",
					dstTimeRange.String(), err)
			}
			xlog.Debugf("[tss] Aggregation job %s completed", dstTimeRange.String())

		case <-ticker30D.C:
			dstTimeRange := tss.TimeRange_TTL_30D
			if err := tsdb.runAggregation(dstTimeRange); err != nil {
				xlog.Errorf("[tss] Unable to complete aggregation job %s: %v",
					dstTimeRange.String(), err)
			}
			xlog.Debugf("[tss] Aggregation job %s completed", dstTimeRange.String())

		case <-ticker365D.C:
			dstTimeRange := tss.TimeRange_TTL_365D
			if err := tsdb.runAggregation(dstTimeRange); err != nil {
				xlog.Errorf("[tss] Unable to complete aggregation job %s: %v",
					dstTimeRange.String(), err)
			}
			xlog.Debugf("[tss] Aggregation job %s completed", dstTimeRange.String())

		case <-tsdb.aggControllerCloseCh:
			return
		}
	}
}

type aggDataPoint struct {
	tmStart  time.Time
	tmStop   time.Time
	metric   tss.MetricType
	valueSum float64
	numKeys  int
}

func (tsdb *tsDB) runAggregation(dstTimeRange tss.TimeRange) error {
	srcTimeRange := aggTaskMap[dstTimeRange].srcTimeRange
	aggInterval := aggTaskMap[dstTimeRange].aggInterval

	srcKeyPrefix := []byte(fmt.Sprintf("%d:", int(srcTimeRange)))
	dstKeyPrefix := []byte(fmt.Sprintf("%d:", int(dstTimeRange)))

	dpLast, err := tsdb.Last(dstKeyPrefix)
	if err != nil {
		return errors.Wrapf(err, "[%v] function tsdb.Last()", errors.Trace())
	}

	dps, err := tsdb.Scan(srcKeyPrefix)
	if err != nil {
		return errors.Wrapf(err, "[%v] function tsdb.Scan()", errors.Trace())
	}

	adpList := make([]*tss.DataPoint, 0)
	adpMap := make(map[tss.MetricType]*aggDataPoint)

	for _, dp := range dps {
		tm := time.UnixMilli(dp.Timestamp)

		if time.Since(tm) < aggInterval {
			continue
		}

		if dpLast != nil {
			tmLast := time.UnixMilli(dpLast.Timestamp)
			if tm.Before(tmLast) {
				continue
			}
		}

		adp, ok := adpMap[dp.Metric]
		if ok {
			if tm.After(adp.tmStart) && tm.Before(adp.tmStop) {
				adp.valueSum += dp.Value
				adp.numKeys++
			} else {
				if tm.After(adp.tmStop) {
					adpList = append(adpList, &tss.DataPoint{
						Timestamp: adp.tmStop.UnixMilli(),
						TimeRange: dstTimeRange,
						Metric:    adp.metric,
						Value:     adp.valueSum / float64(adp.numKeys),
					})
					adpMap[dp.Metric] = newAggDataPoint(dp, aggInterval)
				}
			}
		} else {
			adpMap[dp.Metric] = newAggDataPoint(dp, aggInterval)
		}
	}

	if err := tsdb.WriteBatch(adpList); err != nil {
		return errors.Wrapf(err, "[%v] function tsdb.WriteBatch()", errors.Trace())
	}

	return nil
}

func newAggDataPoint(dp *tss.DataPoint, aggInterval time.Duration) *aggDataPoint {
	adp := &aggDataPoint{
		metric:   dp.Metric,
		valueSum: dp.Value,
		numKeys:  1,
	}
	tmStart := time.UnixMilli(dp.Timestamp)
	tmStop := tmStart.Add(aggInterval)
	adp.tmStart = tmStart
	adp.tmStop = tmStop

	return adp
}
