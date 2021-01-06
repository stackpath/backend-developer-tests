package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FizzBuzzSuite struct {
	suite.Suite
}

func TestFizzBuzzSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzSuite))
}

func (suite FizzBuzzSuite) TestFizzFuzz() {
	cases := []struct {
		total          int64
		fizzAt         int64
		fizzIndex      int
		buzzAt         int64
		buzzIndex      int
		expect         []string
		shouldFizzBuzz bool
		noElem         bool
	}{
		{
			total:     5,
			fizzAt:    3,
			fizzIndex: 2,
			buzzAt:    5,
			buzzIndex: 4,
		},
		{
			total:     100,
			fizzAt:    4,
			fizzIndex: 39,
			buzzAt:    9,
			buzzIndex: 80,
		},
		{
			total:          4,
			fizzAt:         1,
			fizzIndex:      1,
			buzzAt:         1,
			buzzIndex:      1,
			shouldFizzBuzz: true,
		},
		{
			total:     0,
			fizzAt:    0,
			fizzIndex: 0,
			buzzAt:    0,
			buzzIndex: 0,
			noElem:    true,
		},
	}

	for i, c := range cases {
		result := FizzBuzz(c.total, c.fizzAt, c.buzzAt)

		if c.noElem {
			suite.Equal(len(result), 0)
		} else {
			if c.shouldFizzBuzz {
				suite.Equal(result[c.fizzIndex], "FizzBuzz", "Test case #%d", i)
				suite.Equal(result[c.buzzIndex], "FizzBuzz", "Test case #%d", i)
			} else {
				suite.Equal(result[c.fizzIndex], "Fizz", "Test case #%d", i)
				suite.Equal(result[c.buzzIndex], "Buzz", "Test case #%d", i)
			}
		}
		suite.Equal(int64(len(result)), c.total, "Test case #%d", i)

	}
}
