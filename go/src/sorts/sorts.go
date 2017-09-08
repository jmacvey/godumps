package main

import (
	"fmt"
)

func swap(arr []int, indexOne int, indexTwo int) {
	tmp := arr[indexOne]
	arr[indexOne] = arr[indexTwo]
	arr[indexTwo] = tmp
}

func bubbleSort(arr []int) []int {
	offset := 1
	length := len(arr)
	for offset != length { // N executions
		upTo := length - offset
		for i := 0; i < upTo; i++ { // per iteration: (N-1) + (N-2) + (N - 3) + .... + 1 + 0
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1) // O(1) <- 3 statements
			}
		}
		offset++
	}
	// Big(O) ; what does (N - 1) + (N -2) + (N - 3) + ... + 1 + 0 = ?
	// = N + (N - 1) + (N - 2) + ... + 1 + 0 - N == sum(i = 0, N, i) - N
	return arr
}

func main() {
	testArr := []int{10, 8, 1, 2, 4}
	fmt.Println(insertionSort(testArr))
}
