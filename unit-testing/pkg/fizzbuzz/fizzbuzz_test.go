package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		total  int64
		fizzAt int64
		buzzAt int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"success test", args{6, 2, 3}, []string{"1", "Fizz", "Buzz", "Fizz", "5", "FizzBuzz"}},
		{"noFizzBuzz test", args{6, 7, 8}, []string{"1", "2", "3", "4", "5", "6"}},
		{"FullFizzBuzz test", args{4, 1, 3}, []string{"Fizz", "Fizz", "FizzBuzz", "Fizz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.total, tt.args.fizzAt, tt.args.buzzAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
