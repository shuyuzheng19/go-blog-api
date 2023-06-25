package response

import (
	"fmt"
	"time"
)

func FormatTime(date time.Time) string {
	return date.Format("2006年01月02日 15:04")
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func FormatTimeAgo(date time.Time) string {

	second := time.Now().Unix() - date.Unix()

	var dateStr string

	if second <= 60 {
		dateStr = "刚刚"
	} else if second > 60 && second <= 60*60 {
		dateStr = fmt.Sprintf("%d分钟前", second/60)
	} else if second > 60*60 && second <= 60*60*24 {
		dateStr = fmt.Sprintf("%d小时前", second/60/60)
	} else if second > 60*60*24 && second <= 60*60*24*30 {
		dateStr = fmt.Sprintf("%d天前", second/60/60/24)
	} else if second > 60*60*24*30 && second <= 60*60*24*30*12 {
		dateStr = fmt.Sprintf("%d月前", second/60/60/24/30)
	} else {
		dateStr = fmt.Sprintf("%d年前", second/60/60/24/(30*12))
	}

	return dateStr
}
