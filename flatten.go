package fns

func Flatten[T any](gen Generator[[]T]) Generator[T] {
	var buf []T
	done := false
	return func() (item T, err error) {
		for {
			if done && len(buf) == 0 {
				err = &GeneratorDoneError{}
				return
			}
			if len(buf) == 0 {
				buf, err = gen()
				if IsGeneratorDoneError(err) {
					buf = nil
					done = true
					err = nil
					continue
				}
				if err != nil {
					return
				}
			}
			item = buf[0]
			buf = buf[1:]
			return
		}
	}
}
