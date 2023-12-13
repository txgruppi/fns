package fns

func FromSlice[T any](slice []T) Generator[T] {
	curr := 0
	return func() (item T, err error) {
		if curr >= len(slice) {
			err = &GeneratorDoneError{}
			return
		}
		item = slice[curr]
		curr++
		return
	}
}
