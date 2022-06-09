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

// Format ...
func Format(t time.Time, format string) string {
	return t.Format(format)
}
