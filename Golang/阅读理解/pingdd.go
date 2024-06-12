package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := 0
	fmt.Scan(&a)
	arr := make([]string, a)
	var input string
	fmt.Scan(&input)
	for i := range input {
		arr = append(arr, string(input[i]))
	}

	var charMap = map[string]int{"A": 1, "B": 2, "c": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10,
		"K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21, "V": 22,
		"W": 23, "X": 24, "Y": 25, "Z": 26}
	var windows = []string{"P", "D", "D"}
	//get mex value:
	maxValueNum := a / 3
	flag := len(arr)%3 + 1
	maxValueArr := make([]int, flag)
	for j := 0; j < flag; j++ {
		resultValue := 0
		for k := j; k < maxValueNum; k += 3 {
			resultValue = int(math.Abs(float64(charMap[arr[k]]-charMap[windows[0]]) + float64(charMap[arr[k+1]]-charMap[windows[1]]) + float64(charMap[arr[k+2]]-charMap[windows[2]])))
			maxValueArr = append(maxValueArr, resultValue)
		}

	}

	sort.Ints(maxValueArr)
	fmt.Print(maxValueArr)
	//var result = []int{maxValueNum, maxValueArr[0]}

}
