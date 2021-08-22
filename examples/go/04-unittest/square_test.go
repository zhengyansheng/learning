package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	// Table-Driven
	tests := []struct {
		name  string   // uniq name
		input struct { // input
			x int
			y int
		}
		expected int // expected result
	}{
		{
			name: "d1",
			input: struct {
				x int
				y int
			}{1, 2},
			expected: 2,
		},
		{
			name: "d1",
			input: struct {
				x int
				y int
			}{2, 7},
			expected: 14,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Square(test.input.x, test.input.y)
			assert.Equal(t, test.expected, actual)
		})
	}
}
