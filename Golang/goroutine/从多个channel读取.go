/*
*@Author: Wenqiang
*@Date:   2024/4/19
*@Time:   16:53
 */
package main

import (
	"fmt"
	"time"
)

func send(ch chan int, begin int) {
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 10)
	go send(ch2, 0)
	time.Sleep(time.Second)

	for {
		select {
		case val := <-ch1:
			fmt.Println("chan1 : ", val)
		case val := <-ch2:
			fmt.Println("chan2; ", val)
		case <-time.After(2 * time.Second):
			fmt.Println("Time out")
			return
		}

	}
}
