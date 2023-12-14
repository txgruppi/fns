**wip**

A set of Go functions to work with lazy sequences with no reflection and no goroutines.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/txgruppi/fns"
)

type indexedLine struct {
	l string
	i int
}

func run() error {
	reader := strings.NewReader("some\nlines\nof\ntext\n")

	readerGen := fns.FromReader(reader, 16)
	stringGen := fns.Map[[]byte, string](readerGen, func(b []byte) (string, error) {
		return string(b), nil
	})
	stringGen = fns.SplitLinesString(stringGen)
	stringGen = fns.Filter[string](stringGen, func(s string) (bool, error) {
		return len(s) > 0, nil
	})
	linesWithIndex := fns.Fold[string, []indexedLine](stringGen, []indexedLine{}, func(vs []indexedLine, v string) ([]indexedLine, error) {
		vs = append(vs, indexedLine{l: v, i: len(vs)})
		return vs, nil
	})
	lineWithIndex := fns.Flatten[indexedLine](linesWithIndex)
	lineWithIndex = fns.Filter[indexedLine](lineWithIndex, func(s indexedLine) (bool, error) {
		return s.i%2 == 0, nil
	})
	stringGen = fns.Map[indexedLine, string](lineWithIndex, func(s indexedLine) (string, error) {
		return s.l, nil
	})
	stringSlice, err := fns.ToSlice[string](stringGen)()
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", stringSlice) // [some of]
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

```