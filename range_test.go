package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestRange(t *testing.T) {
	cases := []struct {
		name     string
		min      int
		max      int
		step     int
		expected []int
	}{
		{
			name:     "empty",
			min:      0,
			max:      0,
			step:     1,
			expected: []int{},
		},
		{
			name:     "one",
			min:      0,
			max:      1,
			step:     1,
			expected: []int{0},
		},
		{
			name:     "many",
			min:      0,
			max:      10,
			step:     1,
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "step",
			min:      0,
			max:      10,
			step:     2,
			expected: []int{0, 2, 4, 6, 8},
		},
		{
			name:     "negative",
			min:      -10,
			max:      0,
			step:     1,
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
		},
		{
			name:     "reverse",
			min:      10,
			max:      0,
			step:     -1,
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.Range[int](c.min, c.max, c.step)
			i := 0
			for ; true; i++ {
				item, err := actual()
				if fns.IsGeneratorDoneError(err) {
					break
				}
				if err != nil {
					t.Errorf("expected nil, got %v", err)
				}
				if item != c.expected[i] {
					t.Errorf("expected %v, got %v", c.expected[i], item)
				}
			}
			if i != len(c.expected) {
				t.Errorf("expected %v, got %v", len(c.expected), i)
			}
		})
	}
}
