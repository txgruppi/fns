package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestFoldEmpty(t *testing.T) {
	gen := fns.Fold[int, int](fns.FromRange[int](0, 0, 1), 0, func(acc int, item int) (int, error) {
		return acc + item, nil
	})
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item != 0 {
		t.Fatal("expected 0")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestFoldOne(t *testing.T) {
	gen := fns.Fold[int, int](fns.FromRange[int](1, 2, 1), 0, func(acc int, item int) (int, error) {
		return acc + item, nil
	})
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

func TestFoldMany(t *testing.T) {
	gen := fns.Fold[int, int](fns.FromRange[int](1, 5, 1), 0, func(acc int, item int) (int, error) {
		return acc + item, nil
	})
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item != 10 {
		t.Fatal("expected 10")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}
