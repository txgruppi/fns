package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestTake(t *testing.T) {
	gen := func() fns.Generator[int] {
		return fns.Range[int](0, 10, 1)
	}
	cases := []struct {
		name     string
		n        int
		subject  fns.Generator[int]
		expected []int
	}{
		{
			name:     "empty",
			n:        0,
			subject:  gen(),
			expected: []int{},
		},
		{
			name:     "one",
			n:        1,
			subject:  gen(),
			expected: []int{0},
		},
		{
			name:     "some",
			n:        5,
			subject:  gen(),
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "all",
			n:        10,
			subject:  gen(),
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "more",
			n:        15,
			subject:  gen(),
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "negative",
			n:        -1,
			subject:  gen(),
			expected: []int{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.Take[int](c.subject, c.n)
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
