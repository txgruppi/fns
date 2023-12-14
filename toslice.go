package fns

func ToSlice[T any](gen Generator[T]) Generator[[]T] {
	done := false
	return func() ([]T, error) {
		if done {
			return nil, &GeneratorDoneError{}
		}
		result := []T{}
		for {
			item, err := gen()
			if IsGeneratorDoneError(err) {
				done = true
				return result, nil
			}
			if err != nil {
				return nil, err
			}
			result = append(result, item)
		}
	}
}
