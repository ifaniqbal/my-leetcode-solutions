package main

import "fmt"

// 704. Binary Search
// https://leetcode.com/problems/binary-search/
func search(nums []int, target int) int {
	mid := len(nums) / 2
	start := 0
	end := len(nums) - 1
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	for start <= end {
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			if end == mid {
				return -1
			}
			end = mid
			mid = mid - (end-start)/2
		} else {
			if start == mid {
				return -1
			}
			start = mid
			mid = mid + (end-start)/2
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 6))
}
