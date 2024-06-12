/*
*@Author: Wenqiang
*@Date:   2024/4/3
*@Time:   16:10
 */
/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target
写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1
*/
package main

import "fmt"

func main() {
	var nums = []int{-1, 0, 3, 5, 9, 12}
	result := aveFind(nums, 9)
	fmt.Print(result)

}
func aveFind(num []int, target int) int {
	right := len(num) - 1
	left := 0
	for left <= right {
		mid := left + (right-left)/2
		if num[mid] == target {
			return mid
		} else if num[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1

}
