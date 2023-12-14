package fns

func Map[A, B any](gen Generator[A], fn func(A) (B, error)) Generator[B] {
	var zero B
	return func() (B, error) {
		item, err := gen()
		if err != nil {
			return zero, err
		}
		return fn(item)
	}
}
