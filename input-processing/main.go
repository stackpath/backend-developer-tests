package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	// TODO: Look for lines in the STDIN reader that contain "error" and output them.
	ans := []byte{}
	errString := "error"
	errStringPtr := 0
	discard := true

	print := func() {
		if !discard {
			fmt.Println(string(ans))
		} else {
			fmt.Println("discarding", string(ans))
		}
		//resetting the state
		ans = []byte{}
		errStringPtr = 0
		discard = true
	}

	for {
		for errStringPtr < len(errString) {
			readbyte, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					fmt.Println("finished reading the stream")
					print()
					return
				}
				fmt.Println("error reading byte")
				break
			}

			if readbyte == '\n' {
				print()
				break
			}
			ans = append(ans, readbyte)
			if readbyte == errString[errStringPtr] {
				errStringPtr += 1
			} else {
				errStringPtr = 0
				if readbyte == errString[errStringPtr] {
					errStringPtr += 1
				}
				break
			}
		}
		if errStringPtr == len(errString) {
			errStringPtr = 0
			discard = false
		}
	}
}
