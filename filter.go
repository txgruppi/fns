package fns

func Filter[T any](fn func(T) (bool, error), gen Generator[T]) Generator[T] {
	return func() (T, error) {
		for {
			item, err := gen()
			if err != nil {
				return item, err
			}
			ok, err := fn(item)
			if err != nil {
				return item, err
			}
			if ok {
				return item, nil
			}
		}
	}
}
