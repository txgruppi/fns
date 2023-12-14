package fns

func Fold[A, B any](gen Generator[A], curr B, fn func(B, A) (B, error)) Generator[B] {
	return func() (B, error) {
		for {
			item, err := gen()
			if IsGeneratorDoneError(err) {
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
