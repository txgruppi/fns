package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestMapEmpty(t *testing.T) {
	gen := fns.Map[int, int](fns.FromRange[int](0, 0, 1), func(item int) (int, error) {
		return item * 2, nil
	})
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestMapOne(t *testing.T) {
	gen := fns.Map[int, int](fns.FromRange[int](1, 2, 1), func(item int) (int, error) {
		return item * 2, nil
	})
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item != 2 {
		t.Fatal("expected 2")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestMapMany(t *testing.T) {
	gen := fns.Map[int, int](fns.FromRange[int](1, 5, 1), func(item int) (int, error) {
		return item * 2, nil
	})
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
