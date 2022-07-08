package main

import (
	"fmt"
	"time"

	"github.com/stackpath/backend-developer-tests/concurrency"
)

func main() {
	simplePool := concurrency.NewSimplePool(2)
	go func() {
		time.Sleep(1 * 100)
		simplePool.Submit(test)
		simplePool.Submit(test2)
		simplePool.Submit(test)
		simplePool.Submit(test)
		simplePool.Submit(test)
		fmt.Println("All job submitted")
	}()

	<-simplePool.GetAllDone()
	fmt.Println("Shutting down")
}

func test() {
	fmt.Println("Inside the helper method")
	time.Sleep(2 * 100)
}

func test2() {
	fmt.Println("Inside the helper method 2")
	time.Sleep(5 * 100)
}
