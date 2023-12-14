package fns

func Take[T any](gen Generator[T], n int) Generator[T] {
	var zero T
	return func() (T, error) {
		if n <= 0 {
			return zero, &GeneratorDoneError{}
		}
		n--
		return gen()
	}
}
