package fns_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/txgruppi/fns"
)

func TestFromReader(t *testing.T) {
	cases := []struct {
		name       string
		bufferSize int
		r          io.Reader
		expected   [][]byte
	}{
		{
			name:       "empty",
			bufferSize: 3,
			r:          &bytes.Buffer{},
			expected:   [][]byte{},
		},
		{
			name:       "one",
			bufferSize: 5,
			r:          bytes.NewBufferString("h"),
			expected: [][]byte{
				[]byte("h"),
			},
		},
		{
			name:       "many",
			bufferSize: 7,
			r:          bytes.NewBufferString("hello world"),
			expected: [][]byte{
				[]byte("hello w"),
				[]byte("orld"),
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := fns.FromReader(c.r, c.bufferSize)
			i := 0
			for ; true; i++ {
				item, err := actual()
				if fns.IsGeneratorDoneError(err) {
					break
				}
				if !bytes.Equal(item, c.expected[i]) {
					t.Errorf("expected %v, got %v", c.expected[i], item)
				}
			}
			if i != len(c.expected) {
				t.Errorf("expected %v, got %v", len(c.expected), i)
			}
		})
	}
}
