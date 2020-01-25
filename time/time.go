package time

import "time"

func epochFormat(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2, 3:04pm MST")
}

func epochFormatDate(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2 (Monday)")
}

func epochFormatTime(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("3:04pm MST")
}

func epochFormatHour(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	s := epochTime.Format("3pm")
	s = s[:len(s)-1]
	if len(s) == 2 {
		s += " "
	}
	return s
}
