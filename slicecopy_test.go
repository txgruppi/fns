package fns_test

import (
	"bytes"
	"testing"

	"github.com/txgruppi/fns"
)

func TestSliceCopy(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	gen := fns.FromReader(3, buf)
	gen = fns.SliceCopy(gen)
	actual, err := fns.ToSlice[[]byte](gen)()
	if err != nil {
		t.Fatal(err)
	}
	expected := [][]byte{[]byte("hel"), []byte("lo "), []byte("wor"), []byte("ld")}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(actual))
	}
}
