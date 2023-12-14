package fns_test

import (
	"bytes"
	"testing"

	"github.com/txgruppi/fns"
)

func TestFromReaderEmpty(t *testing.T) {
	gen := fns.FromReader(bytes.NewReader([]byte{}), 1)
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestFromReaderOne(t *testing.T) {
	gen := fns.FromReader(bytes.NewReader([]byte{1}), 1)
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item[0] != 1 {
		t.Fatal("expected 1")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestFromReaderMany(t *testing.T) {
	gen := fns.FromReader(bytes.NewReader([]byte{1, 2, 3, 4, 5, 6}), 2)
	for i := 1; i <= 6; i += 2 {
		item, err := gen()
		if err != nil {
			t.Fatal(err)
		}
		if item[0] != byte(i) {
			t.Fatalf("expected %d", i)
		}
		if item[1] != byte(i+1) {
			t.Fatalf("expected %d", i+1)
		}
	}
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}
