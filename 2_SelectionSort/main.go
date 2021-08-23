package main

import "fmt"

func main() {
	a := []int{10, 45, 6, 90, 3, 5, 55, 1, 2, 4}
	fmt.Println("Before", a)
	SelectionSort(&a)
	fmt.Println("After", a)
}

func SelectionSort(arr *[]int) {

	for wall := len(*arr)-1; wall >0; wall-- {
		largest := 0
		for i := 1; i <= wall; i++ {
			if (*arr)[i] > (*arr)[largest] {
				largest = i
			}
		}
		(*arr)[wall], (*arr)[largest] = (*arr)[largest], (*arr)[wall]
	}
}
