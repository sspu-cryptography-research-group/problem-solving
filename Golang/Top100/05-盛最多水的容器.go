package main

import "fmt"

// 超出时间限制
func maxArea1(height []int) int {
	maxWater := 0
	currentWater := 0

	for i := 0; i < len(height); i++ {
		for j := height[i]; j > 0; j-- {
			for k := i + 1; k < len(height); k++ {
				if height[k] < j {
					continue
				}

				currentWater = j * (k - i)
				if currentWater > maxWater {
					maxWater = currentWater
				}
			}
		}
	}
	return maxWater
}

// 题解：双指针
func maxArea2(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Print(maxArea2(nums))
}
