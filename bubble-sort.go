package main

import "fmt"

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				aux := array[j]
				array[j] = array[i]
				array[i] = aux
			}
		}
	}
}

func main() {
	array := []int{44, 2, 200, 1}

	bubbleSort(array)
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
