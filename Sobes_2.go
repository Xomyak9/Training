package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(gen_numbers(30))
}

func gen_numbers(x int) []int64 {
	rand.Seed(time.Now().UnixNano())
	result := make([]int64, 0)
	uniqueElem := map[int64]struct{}{}
	for i := 0; i < x; i++ {
		n := rand.Int63()
		if _, ok := uniqueElem[n]; !ok {
			uniqueElem[n] = struct{}{}
			result = append(result, n)
		}
	}
	return result
}
