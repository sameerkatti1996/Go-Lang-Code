package main

import "fmt"

func partition(data []int, start, end int) (int, []int) {
	pivot := data[start]
	i := start
	j := end + 1

	for i < j {

		i++
		for i < end && data[i] < pivot {
			i++
		}

		j--
		for j > start && data[j] > pivot {
			j--
		}

		data[i], data[j] = data[j], data[i]

	}

	data[i], data[j] = data[j], data[i]
	data[start], data[j] = data[j], data[start]

	return j, data
}

func quickSort(data []int, start, end int) []int {
	if start < end {
		pivot, data := partition(data, start, end)
		data = quickSort(data, start, pivot-1)
		data = quickSort(data, pivot+1, end)
	}
	return data
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(quickSort(a, 0, len(a)-1))
}
