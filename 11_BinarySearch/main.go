package main

import "fmt"

func main() {
	a := []int{-100, -50, -40, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(BinarySearch(-1990, &a))
}

func BinarySearch(num int, arr *[]int) int {
	low := 0
	end := len(*arr) - 1
	for low <= end {
		mid := (low + end) / 2
		if (*arr)[mid] == num {
			return mid
		} else if (*arr)[mid] > num {
			end = mid - 1
		} else if (*arr)[mid] < num {
			low = mid + 1
		}
	}
	return -1
}
