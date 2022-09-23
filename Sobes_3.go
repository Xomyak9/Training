package main

import (
	"fmt"
)

func main() {
	var i int = 0
	cache_arr := make([]int, 0)
	for i != 3 {
		fmt.Println("Your cache: ", cache_arr)
		fmt.Println("1.Add \n2.Delete \n3.Exit")
		var x int
		fmt.Println("Enter the operation number: ")
		fmt.Scanln(&x)
		if x == 1 {
			var y int
			cache_arr, y = set(cache_arr)
			// y := set()
			// cache_arr = append(cache_arr, y)
			fmt.Println("Add:", y)
		} else if x == 2 {
			var y int
			cache_arr, y = delete_val(cache_arr)
			fmt.Println("Removing the", y, "element")
		} else if x == 3 {
			fmt.Println("Exit")
		}
		i = x
	}
}

func set(x []int) ([]int, int) {
	var number int
	fmt.Println("Enter the number you want to add: ")
	fmt.Scanln(&number)
	x = append(x, number)
	return x, number

}

func delete_val(x []int) ([]int, int) {
	var number int
	fmt.Println("Enter the number you want to add: ")
	fmt.Scanln(&number)
	indX := number - 1
	x = append(x[:indX], x[indX+1:]...)
	return x, number
}
