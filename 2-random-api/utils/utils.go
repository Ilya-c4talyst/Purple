package utils

import (
	"math/rand/v2"
	"strconv"
)

// Рандомное число 0-6
func GetRandomNumber() string {

	number := rand.IntN(6) + 1
	numberStr := strconv.Itoa(number)

	return numberStr
}
