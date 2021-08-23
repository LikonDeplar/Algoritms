package main

import (
	"fmt"
)

func main() {
	a := []int{10, 45, 6, 90, 3, 5, 55, 1, 2, 4}
	fmt.Println("Before", a)
	BobbleSort(&a)
	fmt.Println("After", a)
}

func BobbleSortWithWall(arr *[]int) {
	for wall := len(*arr) - 1; wall > 0; wall-- {
		for i := 0; i < wall; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
			}
		}
	}
}

//second variant
func BobbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j :=0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}