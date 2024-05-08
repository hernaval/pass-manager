package utils

import "testing"

func TestFilter(t *testing.T) {
	// given number R
	number := []int{34, 54, 57, 75, -5, 5}

	// when filtering negative number
	negFunc := func(n int) bool { return n < 0 }
	negativeN := Filter(number, negFunc)

	// then
	if len(negativeN) > 1 {
		t.Errorf("There is only one negative number here")
	}

	if negativeN[0] != -5 {
		t.Errorf("The func don't return the negative")
	}
}
