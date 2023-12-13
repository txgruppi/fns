package fns_test

import (
	"strings"
	"testing"

	"github.com/txgruppi/fns"
)

func TestSplit(t *testing.T) {
	gen := fns.FromReader(3, strings.NewReader("some\nlines\nof\ntext\n"))
	toString := fns.Map[[]byte, string](func(item []byte) (string, error) {
		return string(item), nil
	}, gen)
	split := fns.SplitStringLines(toString)
	actual, err := fns.ToSlice[string](split)()
	if err != nil {
		t.Fatal(err)
	}
	expected := []string{"some", "lines", "of", "text"}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(actual))
	}
	for i, item := range actual {
		if item != expected[i] {
			t.Fatalf("expected %q, got %q", expected[i], item)
		}
	}
}
