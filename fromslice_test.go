package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestFromSliceEmpty(t *testing.T) {
	gen := fns.FromSlice[int]([]int{})
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestFromSliceOne(t *testing.T) {
	gen := fns.FromSlice[int]([]int{1})
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

func TestFromSliceMany(t *testing.T) {
	gen := fns.FromSlice[int]([]int{1, 2, 3})
	for i := 1; i <= 3; i++ {
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
