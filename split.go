package fns

import (
	"bytes"
	"strings"
)

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

func SeparatorSplitter[T any](
	contains func(T) bool,
	index func(T) int,
	cut func(T, int) (T, T),
	builder func(T),
	build func() T,
	empty func(T) bool,
	gen Generator[T],
) Generator[T] {
	var zero T
	return Split[T](func(prev []T, generatorDone bool) ([]T, T, error) {
		lineBreak := -1
		for i, chunk := range prev {
			if contains(chunk) {
				lineBreak = i
				break
			}
		}
		if lineBreak != -1 {
			for i := 0; i <= lineBreak; i++ {
				idx := index(prev[0])
				if idx == -1 {
					builder(prev[0])
					prev = prev[1:]
					continue
				}
				left, right := cut(prev[0], idx)
				builder(left)
				prev[0] = right
			}
			return prev, build(), nil
		}
		if !generatorDone {
			return prev, zero, &RequestNextError{}
		}
		if len(prev) == 0 {
			return nil, zero, &GeneratorDoneError{}
		}
		for len(prev) > 0 {
			builder(prev[0])
			prev = prev[1:]
		}
		result := build()
		if empty(result) {
			return nil, zero, &GeneratorDoneError{}
		}
		return nil, result, nil
	}, gen)
}

func SplitLinesString(gen Generator[string]) Generator[string] {
	var builder strings.Builder
	return SeparatorSplitter[string](
		func(v string) bool { return strings.Contains(v, "\n") },
		func(v string) int { return strings.Index(v, "\n") },
		func(v string, idx int) (string, string) { return v[:idx], v[idx+1:] },
		func(v string) { builder.WriteString(v) },
		func() string { result := builder.String(); builder.Reset(); return result },
		func(v string) bool { return v == "" },
		gen,
	)
}

func SplitLinesBytes(gen Generator[[]byte]) Generator[[]byte] {
	var buf bytes.Buffer
	return SeparatorSplitter[[]byte](
		func(v []byte) bool { return bytes.Contains(v, []byte("\n")) },
		func(v []byte) int { return bytes.Index(v, []byte("\n")) },
		func(v []byte, idx int) ([]byte, []byte) { return v[:idx], v[idx+1:] },
		func(v []byte) { buf.Write(v) },
		func() []byte {
			cp := make([]byte, buf.Len())
			copy(cp, buf.Bytes())
			buf.Reset()
			return cp
		},
		func(v []byte) bool { return len(v) == 0 },
		gen,
	)
}
