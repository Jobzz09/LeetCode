package main

import "fmt"

func moveZeroes(nums []int) {
	zeroCount := 0
	for i := 0; i < len(nums)-zeroCount; i++ {
		for nums[i] == 0 && zeroCount < len(nums) {
			k := i
			for (k + 1) < len(nums) {
				nums[k] = nums[k+1]
				k++
			}
			nums[k] = 0
			zeroCount++
		}
	}
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
