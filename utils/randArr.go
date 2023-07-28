package utils

import (
	"math/rand"
	"time"
)

func RandArr(length int, maximum int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maximum)
	}
	return arr
}
