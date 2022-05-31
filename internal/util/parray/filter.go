package parray

import "github.com/thoas/go-funk"

// Filter ...
func Filter(arr interface{}, predicate interface{}) interface{} {
	return funk.Filter(arr, predicate)
}
