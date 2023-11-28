package main

import (
	"fmt"

	"fast/user"
)

// test
func search() int {
	var nums = []int{1, 2, 3, 5, 6, 7}
	var target int = 5

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("hello")
	return -1

}
func main() {
	num := search()
	fmt.Println(num)

	user.User()
}
