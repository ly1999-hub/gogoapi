package pstring

import "encoding/base64"

// Base64EncodeToString ...
func Base64EncodeToString(text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	return encoded
}

// Base64DecodeToString ...
func Base64DecodeToString(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	return string(decoded), err
}
