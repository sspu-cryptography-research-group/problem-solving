/*
*@Author: Wenqiang
*@Date:   2023/4/22
*@Time:   16:27
 */
package main

func longestArithSeqLength(nums []int) int {

	selectSort(nums)
	count := 1
	array := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			t := nums[j] - nums[i]
			for k := j + 1; k < len(nums); k++ {
				if nums[k]-nums[j] == t {
					count++
					break
				}
			}
			array = append(array, count)
		}
	}
	selectSort(array)
	return array[len(array)-1]
}
func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}
func main() {
	////使用标准输入来构建数组
	//var n int
	//fmt.Print("Please input the number of array element:")
	//fmt.Scan(&n)
	//
	//arr := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Printf("Please input the %d element:", i+1)
	//	fmt.Scan(&arr[i])
	//}
	//fmt.Println(longestArithSeqLength(arr))
}
