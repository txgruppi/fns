package fns_test

import (
	"reflect"
	"testing"

	"github.com/txgruppi/fns"
)

func TestSliceCopyEmpty(t *testing.T) {
	gen := fns.SliceCopy[int](fns.FromSlice[[]int]([][]int{}))
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSliceCopyOne(t *testing.T) {
	gen := fns.SliceCopy[int](fns.FromSlice[[]int]([][]int{{1}}))
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

func TestSliceCopyMany(t *testing.T) {
	gen := fns.SliceCopy[int](fns.FromSlice[[]int]([][]int{{1, 2, 3}, {4, 5, 6}}))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, []int{1, 2, 3}) {
		t.Fatal("expected [1 2 3]")
	}
	item, err = gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, []int{4, 5, 6}) {
		t.Fatal("expected [4 5 6]")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}
