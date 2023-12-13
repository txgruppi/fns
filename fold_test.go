package fns_test

import (
	"fmt"
	"testing"

	"github.com/txgruppi/fns"
)

func TestFold(t *testing.T) {
	cases := []struct {
		name     string
		subject  fns.Generator[int]
		init     int
		fn       func(int, int) (int, error)
		expected int
	}{
		{
			name:    "empty",
			subject: fns.Range[int](0, 0, 1),
			init:    0,
			fn: func(acc int, item int) (int, error) {
				return acc + item, nil
			},
			expected: 0,
		},
		{
			name:    "one",
			subject: fns.Range[int](1, 2, 1),
			init:    0,
			fn: func(acc int, item int) (int, error) {
				return acc + item, nil
			},
			expected: 1,
		},
		{
			name:    "many",
			subject: fns.Range[int](0, 10, 1),
			init:    0,
			fn: func(acc int, item int) (int, error) {
				return acc + item, nil
			},
			expected: 45,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.Fold[int, int](c.init, c.fn, c.subject)
			item, err := actual()
			if err != nil {
				t.Errorf("expected nil, got %v", err)
			}
			if item != c.expected {
				t.Errorf("expected %v, got %v", c.expected, item)
			}
		})
	}

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		expected := fmt.Errorf("error")
		actual := fns.Fold[int, int](0, func(acc int, item int) (int, error) { return 0, expected }, fns.Range[int](0, 1, 1))
		_, err := actual()
		if err != expected {
			t.Errorf("expected %v, got %v", expected, err)
		}
	})
}
