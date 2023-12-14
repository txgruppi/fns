package fns

func Filter[T any](gen Generator[T], fn func(T) (bool, error)) Generator[T] {
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
