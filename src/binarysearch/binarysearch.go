package binarysearch

import (
	"fmt"
)

func IterativeBinarySearch(a []int, n int) int {
	temp := a
	low := 0

	for len(temp) != 0 {
		high := len(temp)
		midIndex := high / 2

		fmt.Println(temp, high, midIndex)

		val := temp[midIndex]
		if val > n {
			temp = temp[:midIndex]
		} else if val < n {
			temp = temp[midIndex+1:]
			low = midIndex + 1
		} else {
			return low + midIndex
		}
	}

	return -1
}

func RecursiveBinarySearch(a []int, n int) int {
	if len(a) == 0 {
		return -1
	}

	mid := len(a) / 2
	val := a[mid]

	fmt.Println(a, len(a), mid)

	if val > n {
		return RecursiveBinarySearch(a[:mid], n)
	} else if val < n && mid+1 < len(a) {
		r := RecursiveBinarySearch(a[mid+1:], n)
		if r != -1 {
			return mid + 1 + r
		}
	} else if n == val {
		return mid
	}

	return -1
}
