/*
*@Author: Wenqiang
*@Date:   2023/5/11
*@Time:   22:52
 */
package main

import "fmt"

func removeElement(nums []int, val int) []int {
	length := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 0
			length -= 1
		}
	}
	return nums
}

func main() {
	num := []int{3, 2, 2, 3}
	fmt.Print(removeElement(num, 3))
}
