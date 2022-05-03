package matrixstore

import (
	"fmt"
	"io"
	"strings"
)

var (
	ErrInvalidArgs = fmt.Errorf("invalid args")
)

type MatrixStore[T any] struct {
	array [][]T
	next  <-chan int
}

func NewMatrixStore[T any](rows, columns int) (*MatrixStore[T], error) {

	if rows < 1 || columns < 1 {
		return nil, ErrInvalidArgs
	}

	a := make([][]T, rows)
	for i := 0; i < rows; i++ {
		a[i] = make([]T, columns)
	}

	ch := make(chan int, 1)

	go func() {
		for r := 0; r < rows; r++ {
			for c := 0; c < columns; c++ {
				ch <- r
				ch <- c
			}
		}

		close(ch)
	}()

	return &MatrixStore[T]{a, ch}, nil
}

func (ia *MatrixStore[T]) String() string {
	sb := strings.Builder{}
	for i, r := range ia.array {
		_, err := fmt.Fprintf(&sb, "%d %v\n", i+1, r)
		if err != nil {
			return fmt.Sprintf("%v", err)
		}
	}
	return sb.String()
}

func (ia *MatrixStore[T]) Write(value T) error {

	r, ok := <-ia.next
	if !ok {
		return io.EOF
	}

	c := <-ia.next

	ia.array[r][c] = value

	return nil
}

func (ia *MatrixStore[T]) Full() bool {
	_, ok := <-ia.next
	return !ok
}
