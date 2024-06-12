/*
*@Author: Wenqiang
*@Date:   2024/4/19
*@Time:   16:42
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrintInput(ch chan string) {
	for val := range ch {
		if val == "EOF" {
			break
		}
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan string)
	defer close(ch)

	go PrintInput(ch)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Print("END")
			break
		}
	}

}
