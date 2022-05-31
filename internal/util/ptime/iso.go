package ptime

import "time"

// ParseISODate ...
func ParseISODate(s string) time.Time {
	t, _ := time.Parse(dateLayoutFull, s)
	return t
}
