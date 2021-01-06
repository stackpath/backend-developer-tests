package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	run(os.Stdin)
}

func run(stdin io.Reader) {
	/*
	 Read STDIN into a new buffered reader with 10mb size limit
	*/
	reader := bufio.NewReaderSize(stdin, 1024*1024*10)
	// reader := bufio.NewReader(stdin)

	/*
		2 concurrency is arbitrarily set to demonstrate goroutine buffering to
		reserve control over system resources when input size is arbitrary.
		Enables processing of more than 1 line concurrently, which can improve performance when
		line length is very long.
	*/
	concurrency := 2
	sem := make(chan bool, concurrency)
	exit := make(chan bool, concurrency)
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		sem <- true
		exit <- false
		go readLine(reader, sem, exit, &wg)
		if <-exit {
			// Add waitgroup to esnure all lines are read before exit channel is closed
			wg.Wait()
			return
		}
	}
}

func readLine(reader *bufio.Reader, sem, exit chan bool, wg *sync.WaitGroup) {
	defer func() { <-sem }()
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("An error occurred while reading from STDIN. Error: %v\n", err)
	}

	// Allow exit stream for testing
	if line == "exit\n" {
		exit <- true
	}

	if lineContainsError(&line) {
		fmt.Fprintf(os.Stdout, "\n%s", line)
	}
	wg.Done()
}

func lineContainsError(t *string) bool {
	if t != nil {
		return strings.Contains(*t, "error")
	}
	return false
}
