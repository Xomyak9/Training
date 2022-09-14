package main

import "fmt"

func main() {
	slice_1 := []int{3, 5, 1, 5, 67, 4, 13, 1, 23, 22}
	slice_2 := []int{1, 45, 22, 234, 67, 1900, 5, 3, 4, 234, 56, 45, 47, 32}
	fmt.Println(Check(slice_1, slice_2))
}

func Check(x []int, y []int) []int {
	var numbers []int
	for _, value_1 := range x {
		for _, value_2 := range y {
			if value_1 == value_2 {
				numbers = append(numbers, value_2)
			}
		}
	}
	return arrange(del_rep(numbers))
}

func del_rep(x []int) []int {
	for indX, value := range x {
		for i := 0; i < len(x); i++ {
			if indX == i {
				continue
			}
			if value == x[i] {
				x[i] = x[len(x)-1]
				x[len(x)-1] = 0
				x = x[:len(x)-1]
			}
		}
	}
	return x
}

func arrange(x []int) []int {
	list := []int{}
	l := len(x)
	for i := 0; i < l; i++ {
		min := x[0]
		for _, value_1 := range x {
			if min > value_1 {
				min = value_1
			}
		}
		for ind, value_2 := range x {
			if min == value_2 {
				x = append(x[:ind], x[ind+1:]...)
			}
		}
		list = append(list, min)
	}
	return list
}
