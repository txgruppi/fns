package fns_test

import (
	"reflect"
	"testing"

	"github.com/txgruppi/fns"
)

func TestToSlice(t *testing.T) {
	cases := []struct {
		name     string
		subject  fns.Generator[int]
		expected []int
	}{
		{
			name:     "empty",
			subject:  fns.Range[int](0, 0, 1),
			expected: []int{},
		},
		{
			name:     "one",
			subject:  fns.Range[int](0, 1, 1),
			expected: []int{0},
		},
		{
			name:     "many",
			subject:  fns.Range[int](0, 10, 1),
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := fns.ToSlice[int](c.subject)()
			if err != nil {
				t.Errorf("expected nil, got %v", err)
			}
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("expected %v, got %v", c.expected, actual)
			}
		})
	}
}
