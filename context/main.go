package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker() {
	defer wg.Done()

	// 死循环
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
}

func main() {
	wg.Add(1)
	go worker()

	// 等待 wg 结束
	wg.Wait()

	fmt.Printf("over")
}
