package main

import "fmt"

func insertionSort(data []int, start, end int) []int {
	var key int
	var j int
  
	for i := start; i <= end; i++ {
		j = i
		key = data[i]

		for j > 0 && data[j - 1] > key {
			data[j] = data[j-1]
			j--
		}
		data[j] = key
	}

	return data
}

func main() {
	a := []int{5, 4, 10, 12, 1, 2, 7}
	fmt.Println(insertionSort(a, 0, len(a)-1))
}
