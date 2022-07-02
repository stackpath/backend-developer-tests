package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	// Look for lines in the STDIN reader that contain "error" and output them.
	searchFor := "error"
	_, err := scanLineForString(reader, searchFor)
	if err != nil {
		log.Fatal(err)
	}
}

// scanLineForString reads a single line (i.e. till "/n")
// and checks for "error" string
func scanLineForString(reader *bufio.Reader, searchFor string) ([]byte, error) {
	if reader == nil {
		return nil, fmt.Errorf("error: no reader")
	}
	var lastLine []byte
	for {
		var singleLine, line []byte
		var isPrefix bool = true
		var err error
		for isPrefix && err == nil {
			line, isPrefix, err = reader.ReadLine()
			if err == nil || err == io.EOF {
				singleLine = append(singleLine, line...)
				if err == io.EOF {
					log.Println("Finished reading file")
					return lastLine, nil
				}
			} else if err != nil {
				return nil, fmt.Errorf("error: Failed reading line from input. %s", err)
			}
		}
		if bytes.Contains(singleLine, []byte(searchFor)) {
			println(string(singleLine))
			if !isPrefix {
				lastLine = singleLine
			}
		}
	}
}
