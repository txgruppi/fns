package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestFromSlice(t *testing.T) {
	cases := []struct {
		name     string
		subject  []int
		expected []int
	}{
		{
			name:     "empty",
			subject:  []int{},
			expected: []int{},
		},
		{
			name:     "one",
			subject:  []int{0},
			expected: []int{0},
		},
		{
			name:     "many",
			subject:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.FromSlice[int](c.subject)
			i := 0
			for ; true; i++ {
				item, err := actual()
				if fns.IsGeneratorDoneError(err) {
					break
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
