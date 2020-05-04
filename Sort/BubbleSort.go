package main

import "fmt"

func bubbleSort(data []int, start, end int) []int {
	for i := start; i < end; i++ {
		for j := start; j < end-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func main() {
	a := []int{5, 4, 10, 12, 1, 2, 7}
	fmt.Println(bubbleSort(a, 0, len(a)-1))
}
