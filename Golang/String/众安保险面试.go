package main

import "fmt"

func findCombinations(candidates []int, target int) [][]int {
	var result [][]int
	backtrack(candidates, target, 0, []int{}, &result)
	return result
}

func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
	if target == 0 {
		// 找到一个有效的组合，添加到结果中
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	if target < 0 {
		// 当前的组合不可行
		return
	}

	for i := start; i < len(candidates); i++ {
		// 尝试包含当前的数字
		current = append(current, candidates[i])
		// 递归地尝试剩余的数字
		backtrack(candidates, target-candidates[i], i+1, current, result)
		// 移除当前的数字，以尝试下一个可能的组合
		current = current[:len(current)-1]
	}
}

func main() {
	candidates := []int{1, 2, 3, 4, 5, 6}
	target := 9
	combinations := findCombinations(candidates, target)
	fmt.Println(combinations)
}
