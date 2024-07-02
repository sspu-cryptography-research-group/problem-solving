package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Print(threeSum2(nums))

}

func threeSum(nums []int) [][]int {
	resultSet := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					currentResult := []int{}
					currentResult = append(currentResult, nums[i])
					currentResult = append(currentResult, nums[j])
					currentResult = append(currentResult, nums[k])
					resultSet = append(resultSet, currentResult)
				}

			}
		}
	}
	return removeFunc(resultSet)

}

func removeFunc(slices [][]int) [][]int {
	seen := make(map[string]bool)
	var result [][]int

	for _, slice := range slices {
		hashVal := hash(slice)
		if _, exists := seen[hashVal]; !exists {
			seen[hashVal] = true
			result = append(result, slice)
		}
	}
	return result
}

func hash(s []int) string {
	sort.Ints(s)
	var sb strings.Builder
	for i, v := range s {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

// 题解：
func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
