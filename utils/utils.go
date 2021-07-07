package utils

import (
	"math/rand"
)

func Shortener() string {
	const Bytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, 6)
	for i := range b {
		b[i] = Bytes[rand.Intn(len(Bytes))]
	}
	return string(b)
}
