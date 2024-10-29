package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input         []int
		value         int
		expectedIndex int
	}{
		{[]int{0, 2, 3, 4, 10, 40, 60, 100, 250, 1000}, 60, 6},
		{[]int{-3, -2, -1, 0, 3, 4, 10, 40, 60, 100, 250, 1000}, -1, 2},
		{[]int{-255, -3, -2, -1, 0, 3, 4, 10, 40, 60, 100, 250, 1000}, -255, 0},
		{[]int{-255, -3, -2, -1, 0, 3, 4, 10, 40, 60, 100, 250, 1000}, 1000, 12},
		{[]int{-255, -3, -2, -1, 0, 3, 4, 10, 40, 99, 300, 550, 780}, 120, -1},
		{[]int{}, 2, -1},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedIndex, SearchValueFromSlice(test.input, test.value))
	}
}
