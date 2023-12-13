**wip**

A set of Go functions to work with lazy sequences with no reflection and no goroutines.

```go
package main

import (
	"github.com/txgruppi/fns"
)

func run() error {
	numbers := fns.Range[int](0, 1000, 1)

	odds := fns.Filter[int](func(item int) (bool, error) {
		return item%2 == 1, nil
	}, numbers)

	ten := fns.Take[int](10, odds)

	product := fns.Fold[int, int](1, func(acc int, item int) (int, error) {
		return acc * item, nil
	}, ten)

	result, err := product()
	if err != nil {
		return err
	}

	println(result) // 654729075

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
```