package pstring

import (
	"math/rand"
	"strings"
	"time"

	"myapp/internal/mongodb"
)

const charsetPassword = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetCode = "0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateDefaultPassword ...
func GenerateDefaultPassword(length int) string {
	return stringWithCharset(length, charsetPassword)
}

// GenerateCode ...
func GenerateCode(s string) string {
	var (
		underscore = "_"
		emptySpace = " "
	)
	return strings.ReplaceAll(mongodb.NonAccentVietnamese(s), emptySpace, underscore)
}

// GenerateCodeString ...
func GenerateCodeString(length int) string {
	return stringWithCharset(length, charsetCode)
}
