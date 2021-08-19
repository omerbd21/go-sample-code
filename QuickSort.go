package main

import "fmt"

func QuickSort(arr []int, low int, high int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	if low >= high {
		return nil
	}
	middle := low + (high-low)/2
	pivot := arr[middle]
	i := low
	j := high
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if low < j {
		arr = QuickSort(arr, low, j)
	}
	if high > i {
		arr = QuickSort(arr, i, high)
	}
	return arr
}

func main() {
	arr := []int{67, 34, 256, 1, 56, 980, 12, 7431, 0, 10}
	fmt.Println(QuickSort(arr, 0, len(arr)-1))
}
