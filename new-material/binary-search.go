package main

import "fmt"

func binarySearchNumber(arr []int, target int) int {
	low := 0
	hi := len(arr) - 1
	answer := -1

	for low <= hi {
		mid := low + (hi-low)/2
		if arr[mid] == target {
			answer = mid
			break
		} else if arr[mid] > target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return answer
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("target index at :", binarySearchNumber(arr, 1))
}
