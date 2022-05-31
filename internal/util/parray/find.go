package parray

import "github.com/thoas/go-funk"

// Find ...
func Find(arr interface{}, predicate interface{}) interface{} {
	return funk.Find(arr, predicate)
}

// FindString ...
func FindString(s []string, cb func(s string) bool) (string, bool) {
	return funk.FindString(s, cb)
}
