/*
*@Author: Wenqiang
*@Date:   2024/4/15
*@Time:   19:49
 */
package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	var arr []int
	for i := 0; i < a; i++ {
		x := 0
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	var score = 0
	if !isJumpArray(arr) {
		score = 0
	} else {
		//求arr的所有子数组元素和
		windows := 2
		for windows <= len(arr) {
			for i := 0; i < len(arr)-windows+1; i++ {
				for j := i; j < i+windows; j++ {
					score += arr[j]
				}
			}
			windows += 1

		}
	}
	fmt.Print(score)

}
func isJumpArray(arr []int) bool {
	if len(arr) >= 2 {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] != arr[i+1] {
				continue
			} else {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
