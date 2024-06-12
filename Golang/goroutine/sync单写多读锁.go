/*
*@Author: Wenqiang
*@Date:   2024/4/19
*@Time:   21:03
 */
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var rwLock sync.RWMutex

func main() {

	//获取读锁
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.RLock()
			defer rwLock.RUnlock()
			fmt.Println("Read func" + strconv.Itoa(i) + "get lock at" + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	//
	time.Sleep(time.Second / 10)
	//获取读锁
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.RLock()
			defer rwLock.RUnlock()
			fmt.Println("Write func" + strconv.Itoa(i) + "get lock at" + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(time.Second * 10)
}
