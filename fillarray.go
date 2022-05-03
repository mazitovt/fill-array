package fill_array

import "fmt"

type Writer[T any] interface {
	Write(T) error
	Full() bool
}

type Reader[T any] interface {
	Read() (T, error)
}

var ErrWriterIsNotFull = fmt.Errorf("writer isn't fully filled")

func FullFill[T comparable](w Writer[T], r Reader[T]) (err error) {

	for {
		var v T
		if v, err = r.Read(); err != nil {
			break
		}
		if err = w.Write(v); err != nil {
			break
		}
	}

	if !w.Full() {
		return ErrWriterIsNotFull
	}

	return nil
}
