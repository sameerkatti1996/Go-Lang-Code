package main

import "fmt"

func merge(data []int, start, mid, end int) []int {
	i := start
	j := mid + 1
	temp := []int{}

	for i <= mid && j <= end {
		if data[i] < data[j] {
			temp = append(temp, data[i])
			i++
		} else {
			temp = append(temp, data[j])
			j++
		}
	}

  if i > mid {
		for ;j <= end;j++ {
			temp = append(temp, data[j])
		}
	} else {
		for ;i <= mid;i++ {
			temp = append(temp, data[i])
		}
	}

  for i, j = start, 0;i <= end;i, j = i+1, j+1 {
		data[i] = temp[j]
	}

  return data
}

func mergeSort(data []int, start, end int) []int {
	if start < end {
		mid := int((start + end) / 2)
		data = mergeSort(data, start, mid)
		data = mergeSort(data, mid+1, end)
		data = merge(data, start, mid, end)
	}
	return data
}

func main() {
	a := []int{5, 4, 10, 12, 1, 2, 7}
	fmt.Println(mergeSort(a, 0, len(a)-1))
}
