package fns

func Map[A, B any](fn func(A) (B, error), gen Generator[A]) Generator[B] {
	var zero B
	return func() (B, error) {
		item, err := gen()
		if err != nil {
			return zero, err
		}
		return fn(item)
	}
}
