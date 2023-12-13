package fns

func SliceCopy[T any](gen Generator[[]T]) Generator[[]T] {
	return func() ([]T, error) {
		item, err := gen()
		cp := make([]T, len(item))
		copy(cp, item)
		return cp, err
	}
}
