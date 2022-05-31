package pstring

import "strings"

// ContainsLowercase ...
func ContainsLowercase(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	return strings.Contains(s1, s2)
}
