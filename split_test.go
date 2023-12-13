package fns_test

import (
	"strings"
	"testing"

	"github.com/txgruppi/fns"
)

func TestSplitStringLines(t *testing.T) {
	gen := fns.FromReader(3, strings.NewReader("some\nlines\nof\ntext\n"))
	toString := fns.Map[[]byte, string](func(item []byte) (string, error) {
		return string(item), nil
	}, gen)
	split := fns.SplitLinesString(toString)
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

func TestSplitLinesBytes(t *testing.T) {
	gen := fns.FromReader(4, strings.NewReader("some\nlines\nof\ntext"))
	gen = fns.SliceCopy[byte](gen)
	split := fns.SplitLinesBytes(gen)
	actual, err := fns.ToSlice[[]byte](split)()
	if err != nil {
		t.Fatal(err)
	}
	expected := [][]byte{[]byte("some"), []byte("lines"), []byte("of"), []byte("text")}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(actual))
	}
	for i, item := range actual {
		if string(item) != string(expected[i]) {
			t.Fatalf("expected %q, got %q", expected[i], item)
		}
	}
}
