package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestFilterEmpty(t *testing.T) {
	gen := fns.FromRange[int](0, 0, 1)
	gen = fns.Filter[int](gen, func(item int) (bool, error) {
		return item%2 == 0, nil
	})
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestFilterOne(t *testing.T) {
	gen := fns.FromRange[int](1, 2, 1)
	gen = fns.Filter[int](gen, func(item int) (bool, error) {
		return item%2 == 1, nil
	})
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item != 1 {
		t.Fatal("expected 2")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestFilterMany(t *testing.T) {
	gen := fns.FromRange[int](1, 5, 1)
	gen = fns.Filter[int](gen, func(item int) (bool, error) {
		return item%2 == 1, nil
	})
	for i := 1; i < 5; i += 2 {
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
