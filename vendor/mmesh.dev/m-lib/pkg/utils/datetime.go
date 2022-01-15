package utils

import (
	"strconv"
	"time"

	"mmesh.dev/m-api-go/grpc/common/datetime"
)

func GetDateTime(t *datetime.DateTime) (int64, error) {
	y := int(t.Year)
	m := time.Month(int(t.Month))
	d := int(t.Day)
	hr := int(t.Hour)
	min := int(t.Minute)

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}

	return time.Date(y, m, d, hr, min, 0, 0, loc).Unix(), nil
}

func ParseTimestamp(s string) (time.Time, error) {
	tm, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(tm, 0), nil
}
