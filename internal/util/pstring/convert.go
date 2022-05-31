package pstring

import (
	"encoding/json"
	"strconv"
)

// ToInt ...
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ToFloat64 ...
func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ToString ...
func ToString(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
