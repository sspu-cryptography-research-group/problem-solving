package main

import (
	"fmt"
	"sync"
	"time"
)

func WatiGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("working %d...", id)
	time.Sleep(time.Second)
	fmt.Printf("done %d\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WatiGroup(i, &wg)
	}
	wg.Wait()
	fmt.Println("all done")
}
