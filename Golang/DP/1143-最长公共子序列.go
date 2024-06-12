/*
*@Author: Wenqiang
*@Date:   2024/4/17
*@Time:   10:21
 */
package main

import "fmt"

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = maxValue(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := "abcde"
	b := "ace"
	fmt.Println(longestCommonSubsequence(a, b))
}
