package unit_test

import (
	"testing"
)

func TestSquare(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "d1",
			x:    1,
			y:    2,
		},
		{
			name: "d2",
			x:    2,
			y:    7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := unit_test_test.Square(test.x, test.y)
		})
	}
}
