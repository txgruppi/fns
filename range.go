package fns

import (
	"golang.org/x/exp/constraints"
)

type RangeItem interface {
	constraints.Float | constraints.Integer
}

func FromRange[T RangeItem](min, max, step int) Generator[T] {
	curr := min
	return func() (item T, err error) {
		if (step > 0 && curr >= max) || (step < 0 && curr <= max) {
			err = &GeneratorDoneError{}
			return
		}
		item = T(curr)
		curr += step
		return
	}
}
