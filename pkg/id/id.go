package id

import (
	"math/rand"
	"time"
)

const DEFAULT_LEN = 15
var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")


func GenerateID(prefix string) string{
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, DEFAULT_LEN)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return prefix + string(b)
}