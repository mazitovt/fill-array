package producer

import "io"

type Producer[T comparable] struct {
	unique bool
	count  int64
	used   map[T]struct{}
	next   func() T
}

func NewProducer[T comparable](unique bool, count int64, next func() T) *Producer[T] {

	return &Producer[T]{
		unique,
		count,
		make(map[T]struct{}, count),
		next,
	}
}

func (ur *Producer[T]) Read() (T, error) {

	if !ur.unique {
		return ur.next(), nil
	}

	var v T
	for {
		if ur.count == 0 {
			return v, io.EOF
		}

		v = ur.next()

		if _, ok := ur.used[v]; !ok {
			ur.used[v] = struct{}{}
			ur.count--
			return v, nil
		}
	}
}
