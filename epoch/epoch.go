package epoch

import "time"

// Format : format epoch -- "January 2, 3:04pm MST"
func Format(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2, 3:04pm MST")
}

// FormatDate : format epoch -- "January 2 (Monday)"
func FormatDate(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("January 2 (Monday)")
}

// FormatTime : format epoch -- "3:04pm MST"
func FormatTime(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	return epochTime.Format("3:04pm MST")
}

// FormatHour : format epoch -- "3pm"
func FormatHour(seconds int64) string {
	epochTime := time.Unix(0, seconds*int64(time.Second))
	s := epochTime.Format("3pm")
	s = s[:len(s)-1]
	if len(s) == 2 {
		s += " "
	}
	return s
}
