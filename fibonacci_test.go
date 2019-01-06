package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input uint
		want  uint
	}{
		{input: 0, want: 0},
		{input: 1, want: 1},
		{input: 2, want: 1},
		{input: 3, want: 2},
		{input: 4, want: 3},
		{input: 5, want: 5},
		{input: 6, want: 8},
		{input: 7, want: 13},
		{input: 8, want: 21},
		{input: 9, want: 34},
		{input: 10, want: 55},
	}
	for _, test := range tests {
		if got := fibonacci(test.input); got != test.want {
			t.Errorf("fib(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}
