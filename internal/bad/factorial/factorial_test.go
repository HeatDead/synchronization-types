package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, test := range tests {
		ch := make(chan int)
		go factorial(test.input, ch)
		result := <-ch

		if result != test.expected {
			t.Errorf("Expected factorial of %d to be %d, got %d", test.input, test.expected, result)
		}
	}
}
