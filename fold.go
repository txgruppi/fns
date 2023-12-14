package fns

func Fold[A, B any](gen Generator[A], curr B, fn func(B, A) (B, error)) Generator[B] {
	var zero B
	done := false
	return func() (B, error) {
		if done {
			return zero, &GeneratorDoneError{}
		}
		for {
			item, err := gen()
			if IsGeneratorDoneError(err) {
				done = true
				return curr, nil
			}
			if err != nil {
				return curr, err
			}
			curr, err = fn(curr, item)
			if err != nil {
				return curr, err
			}
		}
	}
}
