package main

import "fmt"

func twoSum(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1

	for start != end {
		sum := nums[start] + nums[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			return []int{start + 1, end + 1}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(twoSum(nums, 6))
}
