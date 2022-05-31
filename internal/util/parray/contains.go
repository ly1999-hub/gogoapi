package parray

import "github.com/thoas/go-funk"

// Contains ...
func Contains(arr interface{}, v interface{}) bool {
	return funk.Contains(arr, v)
}
