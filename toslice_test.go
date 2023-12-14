package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestToSliceEmpty(t *testing.T) {
	gen := fns.ToSlice[int](fns.FromRange[int](0, 0, 1))
	slice, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if len(slice) != 0 {
		t.Fatal("expected empty slice")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestToSliceOne(t *testing.T) {
	gen := fns.ToSlice[int](fns.FromRange[int](1, 2, 1))
	slice, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if len(slice) != 1 {
		t.Fatal("expected slice with one item")
	}
	if slice[0] != 1 {
		t.Fatal("expected 1")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestToSliceMany(t *testing.T) {
	gen := fns.ToSlice[int](fns.FromRange[int](1, 5, 1))
	slice, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if len(slice) != 4 {
		t.Fatal("expected slice with four items")
	}
	for i := 1; i < 5; i++ {
		if slice[i-1] != i {
			t.Fatalf("expected %d", i)
		}
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}
