/*
你有n个小红书账号，每个账号粉丝为a_n，你需要创建一个粉丝为k的新账号，你可以通过老账号发推文来宣传新账号，宣传的方式有两种。
第一种是浅度宣传，浅度宣传可以为新账号增加a_n/2（向下取整）粉丝。
第二种是重度宣传，重度宣传可以为新账号增加a_n粉丝。其中重度宣传你最多使用一次。
求需要最少账号个数。如果不能满足，返回-1
输入：5 8
     1 2 3 4 10
输出：2
*/

package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	var a []int
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		a = append(a, temp)
	}

}
