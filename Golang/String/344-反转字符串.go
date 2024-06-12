/*
*@Author: Wenqiang
*@Date:   2024/4/17
*@Time:   09:49
 */
package main

import "fmt"

// 我的题解
func reverseString(s []string) []string {
	left := 0
	right := len(s) - 1
	if len(s)%2 == 0 {

		for left != right-1 {
			s[left], s[right] = s[right], s[left]
			left += 1
			right -= 1
		}
		return s
	} else {
		for left != right {
			s[left], s[right] = s[right], s[left]
			left += 1
			right -= 1
		}
		return s
	}
	return nil

}

// 官方
func reverseString2(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

func main() {
	var s = []string{"h", "e", "l", "l", "o"}
	fmt.Println(reverseString(s))
}
