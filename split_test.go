package fns_test

import (
	"strings"
	"testing"

	"github.com/txgruppi/fns"
)

func TestSplit(t *testing.T) {
	gen := fns.FromReader(3, strings.NewReader("some\nlines\nof\ntext\n"))
	toString := fns.Map[[]byte, string](func(item []byte) (string, error) {
		return string(item), nil
	}, gen)
	var builder strings.Builder
	split := fns.Split(func(prev []string, generatorDone bool) ([]string, string, error) {
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
			return prev, "", &fns.RequestNextError{}
		}
		for _, chunk := range prev {
			builder.WriteString(chunk)
		}
		return nil, builder.String(), &fns.GeneratorDoneError{}
	}, toString)
	actual, err := fns.ToSlice[string](split)()
	if err != nil {
		t.Fatal(err)
	}
	expected := []string{"some", "lines", "of", "text"}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(actual))
	}
	for i, item := range actual {
		if item != expected[i] {
			t.Fatalf("expected %q, got %q", expected[i], item)
		}
	}
}
