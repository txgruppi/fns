package fns

import "io"

func FromReader(r io.Reader, bufferSize int) Generator[[]byte] {
	buf := make([]byte, bufferSize)
	return func() ([]byte, error) {
		n, err := r.Read(buf)
		if n == 0 || err == io.EOF {
			return nil, &GeneratorDoneError{}
		}
		if err != nil {
			return nil, err
		}
		return buf[:n], nil
	}
}
