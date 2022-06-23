package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name     string
		fizz     int64
		buzz     int64
		total    int64
		expected []string
	}{
		{
			name:     "DifferentFizzBuzz",
			fizz:     1,
			buzz:     2,
			total:    5,
			expected: []string{"Fizz", "FizzBuzz", "Fizz", "FizzBuzz", "Fizz"},
		},
		{
			name:     "SameFizzBuzz",
			fizz:     2,
			buzz:     2,
			total:    5,
			expected: []string{"1", "FizzBuzz", "3", "FizzBuzz", "5"},
		},
		{
			name:     "OnlyFizz",
			fizz:     2,
			buzz:     6,
			total:    5,
			expected: []string{"1", "Fizz", "3", "Fizz", "5"},
		},
		{
			name:     "OnlyBuzz",
			fizz:     6,
			buzz:     2,
			total:    5,
			expected: []string{"1", "Buzz", "3", "Buzz", "5"},
		},
		{
			name:     "NoFizzNoBuzz",
			fizz:     6,
			buzz:     7,
			total:    5,
			expected: []string{"1", "2", "3", "4", "5"},
		},
		{
			name:     "AllFizzBuzz",
			fizz:     1,
			buzz:     1,
			total:    5,
			expected: []string{"FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz"},
		},
		{
			name:     "AllFizz",
			fizz:     1,
			buzz:     6,
			total:    5,
			expected: []string{"Fizz", "Fizz", "Fizz", "Fizz", "Fizz"},
		},
		{
			name:     "AllBuzz",
			fizz:     6,
			buzz:     1,
			total:    5,
			expected: []string{"Buzz", "Buzz", "Buzz", "Buzz", "Buzz"},
		},
		{
			name:     "Fizz_Buzz_FizzBuzz",
			fizz:     2,
			buzz:     3,
			total:    12,
			expected: []string{"1", "Fizz", "Buzz", "Fizz", "5", "FizzBuzz", "7", "Fizz", "Buzz", "Fizz", "11", "FizzBuzz"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := FizzBuzz(c.total, int64(c.fizz), int64(c.buzz))

			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("Expected %v, got %v", c.expected, actual)
			}
		})
	}
}
