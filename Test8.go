package main

import "fmt"

func main() {
	example_1 := map[string]int{
		"number_1": 1,
		"number_2": 2,
		"number_3": 3,
		"number_4": 4,
	}
	example_2 := map[string]int{
		"number_5": 5,
		"number_7": 7,
		"number_9": 9,
	}
	example_3 := map[string]int{
		"number_6": 6,
		"number_8": 8,
	}

	fmt.Println(Merge(Array(example_1, example_2, example_3)))
}

func Array(x ...map[string]int) []map[string]int {
	kol := []map[string]int{}
	kol = append(kol, x...)
	return kol
}

func Merge(x []map[string]int) map[string]int {
	for i := 0; i < len(x)-1; i++ {
		for k, v := range x[i+1] {
			x[0][k] = v
		}
	}
	return x[0]
}
