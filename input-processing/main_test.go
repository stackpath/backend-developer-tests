package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputProcessingTest struct {
	name          string
	searchFor     string
	lineHasString bool
	isError       bool
}

var testCases = []inputProcessingTest{
	{"Happy Path 1", "thingy", true, false},
	{"Happy Path 2", "error", false, false},
	{"Error Path", "", false, true},
}

func TestScanLineForString(t *testing.T) {
	testFile := "test_data.txt"
	file, err := os.Open(testFile)
	defer file.Close()
	if err != nil {
		t.Fatalf("Error reading file: %s", err)
	}
	reader := bufio.NewReader(file)

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				r := reader
				if tc.isError { // forcing error
					r = nil
				}
				line, err := scanLineForString(r, tc.searchFor)
				if tc.lineHasString {
					assert.Contains(t, string(line), tc.searchFor)
				}
				if tc.isError {
					assert.Error(t, err)
				}
			})
	}
}
