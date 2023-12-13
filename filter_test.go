package fns_test

import (
	"fmt"
	"testing"

	"github.com/txgruppi/fns"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		name     string
		subject  fns.Generator[int]
		fn       func(int) (bool, error)
		expected []int
	}{
		{
			name:    "empty",
			subject: fns.Range[int](0, 0, 1),
			fn: func(item int) (bool, error) {
				return true, nil
			},
			expected: []int{},
		},
		{
			name:    "one true",
			subject: fns.Range[int](0, 1, 1),
			fn: func(item int) (bool, error) {
				return true, nil
			},
			expected: []int{0},
		},
		{
			name:    "one false",
			subject: fns.Range[int](0, 1, 1),
			fn: func(item int) (bool, error) {
				return false, nil
			},
			expected: []int{},
		},
		{
			name:    "many even",
			subject: fns.Range[int](0, 10, 1),
			fn: func(item int) (bool, error) {
				return item%2 == 0, nil
			},
			expected: []int{0, 2, 4, 6, 8},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.Filter(c.fn, c.subject)
			i := 0
			for ; true; i++ {
				item, err := actual()
				if fns.IsGeneratorDoneError(err) {
					break
				}
				if err != nil {
					t.Errorf("unexpected error: %v", err)
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

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		expected := fmt.Errorf("error")
		actual := fns.Filter[int](func(item int) (bool, error) { return false, expected }, fns.Range[int](0, 1, 1))
		_, err := actual()
		if err != expected {
			t.Errorf("expected %v, got %v", expected, err)
		}
	})
}
