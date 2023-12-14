package fns_test

import (
	"testing"

	"github.com/txgruppi/fns"
)

func TestFlattenEmpty(t *testing.T) {
	gen := fns.Flatten[int](fns.FromSlice[[]int]([][]int{}))
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestFlattenOne(t *testing.T) {
	gen := fns.Flatten[int](fns.FromSlice[[]int]([][]int{{1}}))
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

func TestFlattenMany(t *testing.T) {
	gen := fns.Flatten[int](fns.FromSlice[[]int]([][]int{{1, 2, 3}, {4, 5, 6}}))
	for i := 1; i <= 6; i++ {
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

func TestFlattenNotUsingValueWithDoneError(t *testing.T) {
	counter := 0
	gen := fns.Flatten[int](func() ([]int, error) {
		if counter == 0 {
			counter++
			return []int{1, 2, 3}, nil
		}
		return []int{4, 5, 6}, &fns.GeneratorDoneError{}
	})
	for i := 0; i < 3; i++ {
		item, err := gen()
		if err != nil {
			t.Fatal(err)
		}
		if item != i+1 {
			t.Fatalf("expected %d", i+1)
		}
	}
	item, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
	if item != 0 {
		t.Fatal("expected zero value")
	}
}
