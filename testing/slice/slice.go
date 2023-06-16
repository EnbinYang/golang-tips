package main

import "fmt"

func processSlice(number []int) {
	fmt.Println(number[0])
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7)

	dst := make([]int, len(numbers))
	copy(dst, numbers)

	fmt.Println(len(dst))
	fmt.Println(cap(dst))

	for index, value := range dst {
		fmt.Println(index, value)
	}

	processSlice(dst)

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, row := range matrix {
		for _, value := range row {
			fmt.Println(value)
		}
	}
}
