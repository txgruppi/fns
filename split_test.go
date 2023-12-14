package fns_test

import (
	"reflect"
	"testing"

	"github.com/txgruppi/fns"
)

func TestSplitLinesStringEmpty(t *testing.T) {
	gen := fns.SplitLinesString(fns.FromSlice[string]([]string{}))
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesStringOneNoLineBreakAtEnd(t *testing.T) {
	gen := fns.SplitLinesString(fns.FromSlice[string]([]string{"some"}))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item != "some" {
		t.Fatal("expected hello")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesStringOneLineBreakAtEnd(t *testing.T) {
	gen := fns.SplitLinesString(fns.FromSlice[string]([]string{"some\n"}))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if item != "some" {
		t.Fatal("expected hello")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesStringManyNoLineBreakAtEnd(t *testing.T) {
	gen := fns.ToSlice(fns.SplitLinesString(fns.FromSlice[string]([]string{"som", "e\nlin", "es\n", "of", "\ntext"})))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, []string{"some", "lines", "of", "text"}) {
		t.Fatal("expected [some lines of text]")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesStringManyLineBreakAtEnd(t *testing.T) {
	gen := fns.ToSlice(fns.SplitLinesString(fns.FromSlice[string]([]string{"som", "e\nlin", "es\n\n", "of", "\ntext\n\n"})))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, []string{"some", "lines", "", "of", "text", ""}) {
		t.Fatal("expected [some lines of text]")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesBytesEmpty(t *testing.T) {
	gen := fns.SplitLinesBytes(fns.FromSlice[[]byte]([][]byte{}))
	_, err := gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesBytesOneNoLineBreakAtEnd(t *testing.T) {
	gen := fns.SplitLinesBytes(fns.FromSlice[[]byte]([][]byte{[]byte("some")}))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, []byte("some")) {
		t.Fatal("expected hello")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesBytesOneLineBreakAtEnd(t *testing.T) {
	gen := fns.SplitLinesBytes(fns.FromSlice[[]byte]([][]byte{[]byte("some\n")}))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, []byte("some")) {
		t.Fatal("expected hello")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesBytesManyNoLineBreakAtEnd(t *testing.T) {
	gen := fns.ToSlice(fns.SplitLinesBytes(fns.FromSlice[[]byte]([][]byte{[]byte("som"), []byte("e\nlin"), []byte("es\n"), []byte("of"), []byte("\ntext")})))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, [][]byte{[]byte("some"), []byte("lines"), []byte("of"), []byte("text")}) {
		t.Fatal("expected [some lines of text]")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}

func TestSplitLinesBytesManyLineBreakAtEnd(t *testing.T) {
	gen := fns.ToSlice(fns.SplitLinesBytes(fns.FromSlice[[]byte]([][]byte{[]byte("som"), []byte("e\nlin"), []byte("es\n\n"), []byte("of"), []byte("\ntext\n\n")})))
	item, err := gen()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(item, [][]byte{[]byte("some"), []byte("lines"), []byte(""), []byte("of"), []byte("text"), []byte("")}) {
		t.Fatal("expected [some lines of text]")
	}
	_, err = gen()
	if !fns.IsGeneratorDoneError(err) {
		t.Fatal("expected GeneratorDoneError")
	}
}
