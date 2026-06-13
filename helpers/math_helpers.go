package helpers

import (
	"math/rand"
	"time"
)

func GenerateGameCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9000) + 1000
}
