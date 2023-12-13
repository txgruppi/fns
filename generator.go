package fns

type GeneratorDoneError struct {
}

func (e *GeneratorDoneError) Error() string {
	return "generator done"
}

func IsGeneratorDoneError(err error) bool {
	_, ok := err.(*GeneratorDoneError)
	return ok
}

type Generator[T any] func() (T, error)
