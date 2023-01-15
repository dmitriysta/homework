package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var first = mergeSort(array[:len(array)/2])
	var second = mergeSort(array[len(array)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	result := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result
}

func main() {
	unsorted := []int{-356, 328, 705, -199, -373, 108, -377, -362, 128, 98, 1, -9, -500, -607, 387, 12, 210, -600, -351, 432}
	sorted := mergeSort(unsorted)
	fmt.Println(sorted)
}
