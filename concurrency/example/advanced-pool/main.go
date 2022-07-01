package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/stackpath/backend-developer-tests/concurrency"
	"golang.org/x/sync/errgroup"
)

func main() {
	maxSlots := 10
	maxConcurrent := 5
	ap, err := concurrency.NewAdvancedPool(maxSlots, maxConcurrent)
	if err != nil {
		log.Fatalf("main: Failed initializing advanced pool. %s", err)
	}

	totalTasks := 20
	resultChan := make(chan string, totalTasks)
	ctx := context.Background()
	group := &errgroup.Group{}
	group.Go(func() error {
		for i := 1; i <= totalTasks; i++ {
			taskID := i
			if err := ap.Submit(ctx, func(ctx context.Context) {
				log.Printf("main: Advanced task %d pending", taskID)
				time.Sleep(time.Second * 2)
				resultChan <- fmt.Sprintf("Advanced task %d finished", taskID)
			}); err != nil {
				resultChan <- fmt.Sprintf("Advanced task %d errored: %s", taskID, err)
			}
		}
		return nil
	})

	for i := 1; i <= totalTasks; i++ {
		result := <-resultChan
		log.Printf("Task result: %s", result)
	}

	if err := group.Wait(); err != nil {
		log.Printf("main: Advanced work failed. %s", err)
	}
	log.Println("All in a days work")
}
