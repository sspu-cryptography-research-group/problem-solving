/*
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/

package main

import (
	"fmt"
	"sort"
)

func groupAngrams(strs []string) [][]string {
	mp := map[string][]string{} //{"key":"[value1,value2...]"}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str) //哈希表的键为一组字母异位词的标志，哈希表的值为一组字母异位词列表。
	}

	ans := make([][]string, 0, len(mp))

	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Print(groupAngrams(strs))
}
