package main

import (
	"fmt"
)

func main() {
	a := []int{-55, -90, 74, 20, 16, 4, 4, -59, 2, 1, 79, 4, 73, 1, 999}
	fmt.Println("Before:", a)
	QuickSort(a)
	fmt.Println("After:", a)

}

func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, low, hig int) {
	if hig-low < 1 {
		return
	}

	pivot := arr[hig]

	index := low

	for i := low; i < hig; i++ {
		if arr[i] < pivot {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index], arr[hig] = arr[hig], arr[index]

	sort(arr, low, index-1)
	sort(arr, index+1, hig)
}
