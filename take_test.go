package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestTakeEmpty(t *testing.T) {
	gen := fns.Take[int](fns.FromRange[int](0, 10, 1), 0)
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestTakeOne(t *testing.T) {
	gen := fns.Take[int](fns.FromRange[int](1, 11, 1), 1)
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item != 1 {
		t.Fatal("expected 1")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestTakeMany(t *testing.T) {
	gen := fns.Take[int](fns.FromRange[int](2, 12, 1), 5)
	for i := 2; i < 7; i++ {
		item, err := gen()
		if err != nil {
			t.Fatal(err)
		}
		if item != i {
			t.Fatalf("expected %d", i)
		}
	}
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}
