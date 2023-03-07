package main

import "fmt"

func concatArraysToSlice(arr1 [3]int, arr2 [5]int) []int {
	result := make([]int, len(arr1)+len(arr2))
	copy(result[:], arr1[:])

	for i, v := range arr2 {
		j := len(arr1) + i
		result[j] = v
	}
	return result
}

func concatArraysToArray(arr1 [3]int, arr2 [5]int) [8]int {
	result := [8]int{}
	copy(result[:], arr1[:])

	for i, v := range arr2 {
		j := len(arr1) + i
		result[j] = v
	}
	return result
}

func concatSlicesToArray(slice1 []int, slice2 []int) {
	data := make([]int, 0)
	data = append(data, slice1...)
	data = append(data, slice2...)
	fmt.Println(data)
}

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [5]int{10, 9, 8, 7, 6}
	fmt.Println(concatArraysToArray(arr1, arr2))
	fmt.Println(concatArraysToSlice(arr1, arr2))
	concatSlicesToArray([]int{1, 2, 3}, []int{4, 5, 6})
}
