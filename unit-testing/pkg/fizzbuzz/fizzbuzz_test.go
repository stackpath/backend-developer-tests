package fizzbuzz

import (
	"fmt"
	"testing"
)

type fizzBuzzTest struct {
	name   string
	total  int64
	fizzAt int64
	buzzAt int64
	result []string
}

var testCases = []fizzBuzzTest{
	{"Negative Total", -20, 0, 0, []string{}},
	{"Negative Fizz", 6, -2, 3, []string{"1", "2", "Buzz", "4", "5", "Buzz"}},
	{"Negative Buzz", 6, 2, -3, []string{"1", "Fizz", "3", "Fizz", "5", "Fizz"}},
	{"Negative Fizz Buzz", 6, -2, -3, []string{"1", "2", "3", "4", "5", "6"}},
	{"All Zeros Input", 0, 0, 0, []string{}},
	{"Zero with Fizz", 0, 1, 0, []string{}},
	{"Zero with Buzz", 0, 0, 1, []string{}},
	{"Fizz Edge", 1, 0, 1, []string{"Buzz"}}, // Fizz panic without > 0 check
	{"Buzz Edge", 1, 1, 0, []string{"Fizz"}}, // Buzz panic without > 0 check
	{"One Fizz Buzz", 1, 1, 1, []string{"FizzBuzz"}},
	{"Happy Path", 10, 2, 3, []string{
		"1", "Fizz", "Buzz", "Fizz", "5",
		"FizzBuzz", "7", "Fizz", "Buzz", "Fizz"},
	},
	{"Happy Path, Defaults", 20, 3, 5, []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
		"16", "17", "Fizz", "19", "Buzz"},
	},
}

func TestFizzBuzz(t *testing.T) {
	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf(
				"Test Case: %s, Total: %d, FizzAt: %d, BuzzAt: %d",
				tc.name,
				tc.total,
				tc.fizzAt,
				tc.buzzAt,
			),
			func(t *testing.T) {
				r := FizzBuzz(tc.total, tc.fizzAt, tc.buzzAt)
				// Check if the total iterations match the provided total
				if tc.total > 0 && len(r) != int(tc.total) {
					t.Errorf("Total, Expected: %d, Got: %d", tc.total, len(r))
				}
				// Check if the result match the expectations
				for i, e := range tc.result {
					if len(r) > 0 && r[i] != e {
						t.Errorf("Expected %s, Got %s", e, r[i])
					}
				}
			},
		)
	}
}
