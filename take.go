package fns

func Take[T any](n int, gen Generator[T]) Generator[T] {
	var zero T
	return func() (T, error) {
		if n <= 0 {
			return zero, &GeneratorDoneError{}
		}
		n--
		return gen()
	}
}
