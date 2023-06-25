package utils

import (
	"time"
)

func Format(timeStamp int64) string {
	return FormatDate(time.Unix(timeStamp, 0))
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func UnixTimeStampToTime(timeStamp int64) time.Time {
	return time.Unix(timeStamp, 0)
}

func MillTimeStampToTime(timeStamp int64) time.Time {
	return time.UnixMilli(timeStamp)
}
