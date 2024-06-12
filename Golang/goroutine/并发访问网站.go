/*
*@Author: Wenqiang
*@Date:   2024/4/17
*@Time:   17:33
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var arr = []string{"http://www.baidu.com", "http://www.google.com", "http://www.bing.com"}
	for _, url := range arr {
		wg.Add(1)
		go fetch(&wg, url)
	}
	wg.Wait()
}

func fetch(wg *sync.WaitGroup, url string) {
	defer wg.Done()

	fmt.Println("url:", url)
}
