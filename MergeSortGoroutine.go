package main

import (
	"fmt"
	"sync"
)

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	var first, second []int

	go func() {
		first = mergeSort(array[:len(array)/2])
		wg.Done()
	}()

	go func() {
		second = mergeSort(array[len(array)/2:])
		wg.Done()
	}()

	wg.Wait()

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
	array := []int{-356, 328, 705, -199, -373, 108, -377, -362, 128, 98, 1, -9, -500, -607, 387, 12, 210, -600, -351, 432}
	sorted := mergeSort(array)
	fmt.Println(sorted)
}
