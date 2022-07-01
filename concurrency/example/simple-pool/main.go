package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stackpath/backend-developer-tests/concurrency"
	"golang.org/x/sync/errgroup"
)

func main() {
	group := &errgroup.Group{}
	totalWorkers := 5
	sp := concurrency.NewSimplePool(totalWorkers)
	sp.Start()

	totalTasks := 20
	resultChan := make(chan string, totalTasks)
	group.Go(func() error {
		for i := 1; i <= totalTasks; i++ {
			taskID := i
			sp.Submit(func() {
				log.Printf("main: Simple task %d pending", taskID)
				time.Sleep(time.Second * 2)
				resultChan <- fmt.Sprintf("Simple task %d finished", taskID)
			})
		}
		return nil
	})

	for i := 1; i <= totalTasks; i++ {
		result := <-resultChan
		log.Printf("Task result: %s", result)
	}

	if err := group.Wait(); err != nil {
		log.Printf("main: Simple work failed. %s", err)
	}
	log.Println("All in a days work")
}
