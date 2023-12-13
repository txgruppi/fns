package fns_test

import (
	"fmt"
	"testing"

	"github.com/txgruppi/fns"
)

func TestMap(t *testing.T) {
	cases := []struct {
		name     string
		subject  fns.Generator[int]
		fn       func(int) (string, error)
		expected []string
	}{
		{
			name:     "empty",
			subject:  fns.Range[int](0, 0, 1),
			fn:       func(item int) (string, error) { return "", nil },
			expected: []string{},
		},
		{
			name:     "one",
			subject:  fns.Range[int](0, 1, 1),
			fn:       func(item int) (string, error) { return "0", nil },
			expected: []string{"0"},
		},
		{
			name:     "many",
			subject:  fns.Range[int](0, 10, 1),
			fn:       func(item int) (string, error) { return fmt.Sprintf("%d", item), nil },
			expected: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.Map[int, string](c.fn, c.subject)
			for i := 0; true; i++ {
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
		})
	}

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		expected := fmt.Errorf("error")
		actual := fns.Map[int, string](func(item int) (string, error) { return "", expected }, fns.Range[int](0, 1, 1))
		_, err := actual()
		if err != expected {
			t.Errorf("expected %v, got %v", expected, err)
		}
	})
}
