package main

import "fmt"

func main() {
	a := []int{10, 45, 6, 90, 3, 5, 55, 1, 2, 4}
	fmt.Println("Before", a)
	InsertionSort(&a)
	fmt.Println("After", a)
}

func InsertionSort(arr *[]int) {

	for wall := 1; wall < len(*arr); wall++ {
		unSorted := (*arr)[wall]
		i := 0
		for i = wall; i > 0 && (*arr)[i-1] > unSorted; i-- {
			(*arr)[i] = (*arr)[i-1]
		}
		(*arr)[i] = unSorted
	}
}