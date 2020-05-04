package main

import (
	"fmt"
)

func selectionSort(data []int, start, end int) []int {
	var indexOfLowest int
	var lowest int

	for i := start; i < end; i++ {

		indexOfLowest = i
		lowest = data[i]

		for j := i + 1; j <= end; j++ {
			if lowest > data[j] {
				lowest = data[j]
				indexOfLowest = j
			}
		}

		data[i], data[indexOfLowest] = data[indexOfLowest], data[i]

	}

	return data
}

func main() {
	a := []int{4, 2, 5, 1, 7, 8}
	a = selectionSort(a, 0, len(a)-1)
	fmt.Println(a)
}
