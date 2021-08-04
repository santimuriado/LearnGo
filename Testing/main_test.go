package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, calculate(2), 4)
}

func TestTableCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
	}
	for _, test := range tests {
		assert.Equal(calculate(test.input), test.expected)
	}
}
