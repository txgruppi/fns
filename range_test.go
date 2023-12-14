package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestRangeEmpty(t *testing.T) {
	gen := fns.FromRange[int](0, 0, 1)
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestRangeOne(t *testing.T) {
	gen := fns.FromRange[int](1, 2, 1)
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

func TestRangeMany(t *testing.T) {
	gen := fns.FromRange[int](2, 10, 2)
	for i := 2; i < 10; i += 2 {
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

func TestRangeReverse(t *testing.T) {
	gen := fns.FromRange[int](10, 2, -2)
	for i := 10; i > 2; i -= 2 {
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
