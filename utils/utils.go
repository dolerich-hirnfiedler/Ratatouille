package utils

import (
	"math/rand"
	"runtime"
)

func RandomString(length int) string {
	// return a random string of length lenght
	result := ""
	for i := 0; i < length; i++ {
		result += string(rune(65 + rand.Intn(25)))
	}
	return result
}

func GetOs() string {
	return runtime.GOOS
}
