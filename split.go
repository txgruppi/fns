package fns

type RequestNextError struct{}

func (e *RequestNextError) Error() string {
	return "request next item"
}

func IsRequestNextError(err error) bool {
	_, ok := err.(*RequestNextError)
	return ok
}

func Split[T any](fn func(prev []T, done bool) (rest []T, result T, err error), gen Generator[T]) Generator[T] {
	chunks := []T{}
	done := false
	return func() (result T, err error) {
		if len(chunks) == 0 && !done {
			err = &RequestNextError{}
		}
		for {
			if err == nil {
				chunks, result, err = fn(chunks, done)
				if IsGeneratorDoneError(err) {
					return
				}
			}
			if IsRequestNextError(err) {
				next, e := gen()
				if done = IsGeneratorDoneError(e); done || e == nil {
					chunks = append(chunks, next)
					err = nil
					continue
				}
				return result, e
			}
			return
		}
	}
}
