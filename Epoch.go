package toolbox

import "time"

func TimeUTC(seconds int, offset int, timezone string, format string) string {
	if format == "" {
		format = "Mon, Jan, 2 - 03:04 PM"
	}
	if timezone == "" {
		timezone = "UTC"
	}
	loc := time.FixedZone(timezone, offset)
	localTime := time.Unix(int64(seconds), 0).In(loc)
	return localTime.Format(format)
}
