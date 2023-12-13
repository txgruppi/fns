package fns

import (
	"golang.org/x/exp/constraints"
)

type RangeItem interface {
	constraints.Float | constraints.Integer
}

func Range[T RangeItem](min, max, step int) Generator[T] {
	curr := min
	return func() (item T, err error) {
		if curr >= max {
			err = &GeneratorDoneError{}
			return
		}
		item = T(curr)
		curr += step
		return
	}
}
