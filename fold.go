package fns

func Fold[A, B any](curr B, fn func(B, A) (B, error), gen Generator[A]) Generator[B] {
	return func() (B, error) {
		for {
			item, err := gen()
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
