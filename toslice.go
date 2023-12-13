package fns

func ToSlice[T any](gen Generator[T]) ([]T, error) {
	result := []T{}
	for {
		item, err := gen()
		if IsGeneratorDoneError(err) {
			return result, nil
		}
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
}
