/*
*@Author: Wenqiang
*@Date:   2024/4/9
*@Time:   11:16
 */
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var left = 0
	var right = left + 1
	result := make([]int, 0)
	for right < len(nums) {
		sum := nums[left] + nums[right]
		if sum == target {
			result = append(result, left)
			result = append(result, right)
			return result
		} else {
			left += 1
			right = left + 1
		}
	}
	return nil
}

func main() {
	var nums = []int{3, 2, 3}
	var target = 6
	fmt.Println(twoSum(nums, target))
}
