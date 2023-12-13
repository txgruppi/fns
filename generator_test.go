package fns_test

import (
	"fmt"
	"testing"

	"github.com/txgruppi/fns"
)

func TestIsGeneratorDoneError(t *testing.T) {
	cases := []struct {
		name     string
		subject  error
		expected bool
	}{
		{
			name:     "nil",
			subject:  nil,
			expected: false,
		},
		{
			name:     "not generator done error",
			subject:  fmt.Errorf("not generator done error"),
			expected: false,
		},
		{
			name:     "generator done error",
			subject:  &fns.GeneratorDoneError{},
			expected: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.IsGeneratorDoneError(c.subject)
			if actual != c.expected {
				t.Errorf("expected %v, got %v", c.expected, actual)
			}
		})
	}
}
