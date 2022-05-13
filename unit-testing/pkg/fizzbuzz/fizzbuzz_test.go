package fizzbuzz

import (
	"strconv"
	"testing"
)

// Test in order to have fizz result for all values
// divisible by fizzAt and not buzzAt
func TestFizzAt(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i += fizzAt {
		if i != 0 && i%fizzAt == 0 && !(i%buzzAt == 0) {
			if result[i-1] != "Fizz" {
				t.Fail()
			}
		}
	}
}

// Test in order to have Buzz result for all values
// divisible by buzzAt and not fizzAt
func TestBuzzAt(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i += buzzAt {
		if i != 0 && i%buzzAt == 0 && !(i%fizzAt == 0) {
			if result[i-1] != "Buzz" {
				t.Fail()
			}
		}
	}
}

// Test in order to have FizzBuzz result for all values
// divisible by fizzAt and buzzAt
func TestFizzBuzzAt(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i += (fizzAt * buzzAt) {
		if i != 0 && i%buzzAt == 0 && i%fizzAt == 0 {
			if result[i-1] != "FizzBuzz" {
				t.Fail()
			}
		}
	}
}

// Test in order to have result string containing Fizz, Buzz or FizzBuzz
// for all values divisible by fizzAt OR buzzAt
func TestFizzBuzzOutput(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i++ {
		if i != 0 && (i%buzzAt == 0 || i%fizzAt == 0) {
			if !((result[i-1] == "Fizz") ||
				(result[i-1] == "Buzz") ||
				(result[i-1] == "FizzBuzz")) {
				t.Fail()
			}
		}
	}
}

// Test in order to have a result number equal to the valued FizzBuzzed
func TestNumberOutput(t *testing.T) {
	total := int64(100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)

	for i := int64(0); i < total; i++ {
		if i != 0 && !(i%buzzAt == 0) && !(i%fizzAt == 0) {
			val, err := strconv.Atoi(result[i-1])
			if err != nil || int64(val) != i {
				t.Fail()
			}
		}
	}
}

// Test with total < zero
func TestNegativeTotal(t *testing.T) {
	total := int64(-100)
	fizzAt := int64(3)
	buzzAt := int64(5)

	result := FizzBuzz(total, fizzAt, buzzAt)
	if len(result) != 1 || result[0] != errorMsg {
		t.Fail()
	}

}

func TestFizzBuzz(t *testing.T) {
	t.Skip("TODO: Add tests")
}
