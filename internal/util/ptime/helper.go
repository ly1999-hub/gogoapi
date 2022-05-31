package ptime

import "time"

//
// NOTE: due to unique timezone in server's code, all using time will be convert to HCM timezone (UTC +7)
// All functions generate time, must be call util functions here
// WARNING: don't accept call time.Now() directly
//

// GetHCMLocation ...
func GetHCMLocation() *time.Location {
	l, _ := time.LoadLocation(TimezoneHCM)
	return l
}

// TimeStartOfDayInHCM ...
func TimeStartOfDayInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	y, m, d := t.In(l).Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, l).UTC()
	return date
}

func TimeOfDayInHCM(t time.Time) time.Time {
	l, _ := time.LoadLocation(TimezoneHCM)
	return t.In(l)
}

// ParseWithLayout ...
func ParseWithLayout(layout, value string) time.Time {
	t, _ := time.Parse(layout, value)
	return t
}

// Format ...
func Format(t time.Time, format string) string {
	return t.Format(format)
}

// TimeFromTimestamp ...
func TimeFromTimestamp(value int64) time.Time {
	return time.Unix(value, 0)
}
