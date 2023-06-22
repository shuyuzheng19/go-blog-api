package utils

import (
	"math/rand"
	"strconv"
)

func CreateRandomNumber() string {
	return strconv.Itoa(rand.Intn(900000) + 100000)
}
