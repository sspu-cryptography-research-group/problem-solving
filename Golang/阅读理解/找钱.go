package main

import (
	"fmt"
)

func main() {
	var value int
	fmt.Scan(&value)

	money := 1024 - value
	if money < 0 {
		return
	}
	var a, b, c, d int
	d = money / 64
	dd := money % 64
	if dd >= 16 {
		c = dd / 16
		cc := dd % 16
		if cc >= 4 {
			b = cc / 4
			a = cc % 4
			fmt.Println(a + b + c + d)
		} else {
			fmt.Println(d + c + cc)
		}
	} else if dd >= 4 {
		c = dd / 4
		cc := dd % 4
		fmt.Println(d + c + cc)
	} else {
		fmt.Println(d + dd)
	}

}
