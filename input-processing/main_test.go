package main

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}

func (suite MainSuite) TestRun() {
	cases := []struct {
		input                string
		shouldErr            bool
		expectedErrOccurance int
	}{
		{
			input:                "error\nexit\n",
			expectedErrOccurance: 1,
		},
		{
			input:                genInput(3),
			expectedErrOccurance: 2,
		},
	}

	for _, c := range cases {
		reader := strings.NewReader(c.input)

		out := readStdOut(func() {
			run(reader)
		})

		lines := strings.Split(out, "\n")
		count := 0
		for _, l := range lines {
			if strings.Contains(l, "error") {
				count += 1
			}
		}
		suite.Equal(c.expectedErrOccurance, count)
	}

}

// Capture std out during the test run to validate outputs
func readStdOut(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()
	os.Stdout = w
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		_, _ = io.Copy(&buf, r)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	w.Close()
	return <-out
}

func genInput(totalLines int) string {
	builder := strings.Builder{}
	for i := 0; i < totalLines; i++ {
		line := genLine()
		builder.WriteString(line)
		if i%2 == 0 {
			builder.WriteString("error\n")
		}
	}
	builder.WriteString("exit\n")
	return builder.String()
}

func genLine() string {
	rand.Seed(time.Now().Unix())
	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length := 1024 * 1024
	var output strings.Builder
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
