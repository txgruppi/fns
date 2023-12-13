package fns

import "strings"

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

func SplitStringLines(gen Generator[string]) Generator[string] {
	var builder strings.Builder
	return Split(func(prev []string, generatorDone bool) ([]string, string, error) {
		lineBreak := -1
		for i, chunk := range prev {
			if strings.Contains(chunk, "\n") {
				lineBreak = i
				break
			}
		}
		builder.Reset()
		if lineBreak != -1 {
			for i := 0; i <= lineBreak; i++ {
				idx := strings.Index(prev[0], "\n")
				if idx == -1 {
					builder.WriteString(prev[0])
					prev = prev[1:]
					continue
				}
				builder.WriteString(prev[0][:idx])
				prev[0] = prev[0][idx+1:]
			}
			return prev, builder.String(), nil
		}
		if !generatorDone {
			return prev, "", &RequestNextError{}
		}
		for _, chunk := range prev {
			builder.WriteString(chunk)
		}
		return nil, builder.String(), &GeneratorDoneError{}
	}, gen)
}
