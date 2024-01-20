package metricsdb

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/nstore"
	"mmesh.dev/m-api-go/grpc/resources/nstore/metricsdb"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

type aggTask struct {
	srcTimeRange nstore.TimeRange
	aggInterval  time.Duration
}

var aggTaskMap map[nstore.TimeRange]*aggTask = map[nstore.TimeRange]*aggTask{
	nstore.TimeRange_TTL_6H: {
		srcTimeRange: nstore.TimeRange_TTL_1H,
		aggInterval:  12 * time.Minute,
	},
	nstore.TimeRange_TTL_12H: {
		srcTimeRange: nstore.TimeRange_TTL_1H,
		aggInterval:  24 * time.Minute,
	},
	nstore.TimeRange_TTL_24H: {
		srcTimeRange: nstore.TimeRange_TTL_6H,
		aggInterval:  48 * time.Minute,
	},
	nstore.TimeRange_TTL_7D: {
		srcTimeRange: nstore.TimeRange_TTL_12H,
		aggInterval:  336 * time.Minute,
	},
	nstore.TimeRange_TTL_14D: {
		srcTimeRange: nstore.TimeRange_TTL_24H,
		aggInterval:  672 * time.Minute,
	},
	nstore.TimeRange_TTL_30D: {
		srcTimeRange: nstore.TimeRange_TTL_7D,
		aggInterval:  1344 * time.Minute,
	},
	nstore.TimeRange_TTL_365D: {
		srcTimeRange: nstore.TimeRange_TTL_30D,
		aggInterval:  12 * 24 * time.Hour,
	},
}

func (tsdb *tsDB) aggController() {
	ticker := time.NewTicker(300 * time.Second) // 5 minutes
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for tr := range aggTaskMap {
				dstTimeRange := tr
				if err := tsdb.runAggregation(dstTimeRange); err != nil {
					xlog.Errorf("[metricsdb] Unable to complete aggregation job %s: %v",
						dstTimeRange.String(), err)
				}
				xlog.Debugf("[metricsdb] Aggregation job %s completed", dstTimeRange.String())
			}

		case <-tsdb.aggControllerCloseCh:
			return
		}
	}
}

type aggDataPoint struct {
	tmStart  time.Time
	tmStop   time.Time
	metric   metricsdb.HostMetricType
	valueSum float64
	numKeys  int
}

func (tsdb *tsDB) runAggregation(dstTimeRange nstore.TimeRange) error {
	srcTimeRange := aggTaskMap[dstTimeRange].srcTimeRange
	aggInterval := aggTaskMap[dstTimeRange].aggInterval

	srcKeyPrefix := []byte(fmt.Sprintf("%s:%d:", hostMetricsPrefix, int(srcTimeRange)))
	dstKeyPrefix := []byte(fmt.Sprintf("%s:%d:", hostMetricsPrefix, int(dstTimeRange)))

	dpLast, err := tsdb.Last(dstKeyPrefix)
	if err != nil {
		return errors.Wrapf(err, "[%v] function tsdb.Last()", errors.Trace())
	}

	dps, err := tsdb.Scan(srcKeyPrefix)
	if err != nil {
		return errors.Wrapf(err, "[%v] function tsdb.Scan()", errors.Trace())
	}

	adpList := make([]*metricsdb.HostMetricDataPoint, 0)
	adpMap := make(map[metricsdb.HostMetricType]*aggDataPoint)

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
					adpList = append(adpList, &metricsdb.HostMetricDataPoint{
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

func newAggDataPoint(dp *metricsdb.HostMetricDataPoint, aggInterval time.Duration) *aggDataPoint {
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
