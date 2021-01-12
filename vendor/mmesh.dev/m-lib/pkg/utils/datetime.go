package utils

import (
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
