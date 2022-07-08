package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	// TODO: Look for lines in the STDIN reader that contain "error" and output them.

	content, err := reader.ReadBytes('\n')

	for err == nil {
		if strings.Contains(string(content), "error") {
			fmt.Println("STDOUT: ", string(content))
		}
		content, err = reader.ReadBytes('\n')
	}
}
