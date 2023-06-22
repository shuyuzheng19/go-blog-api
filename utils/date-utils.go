package utils

import (
	"fmt"
	"time"
)

func Format(timeStamp int64) string {
	return FormatDate(time.Unix(timeStamp, 0))
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func FormatTimeAgo(givenTime time.Time) string {
	duration := time.Since(givenTime)

	if years := int(duration.Hours() / 24 / 365); years > 0 {
		return fmt.Sprintf("%d年前", years)
	} else if months := int(duration.Hours() / 24 / 30); months > 0 {
		return fmt.Sprintf("%d月前", months)
	} else if days := int(duration.Hours() / 24); days > 0 {
		return fmt.Sprintf("%d天前", days)
	} else if hours := int(duration.Hours()); hours > 0 {
		return fmt.Sprintf("%d小时前", hours)
	} else if minutes := int(duration.Minutes()); minutes > 0 {
		return fmt.Sprintf("%d分钟前", minutes)
	} else {
		return "刚刚"
	}
}
