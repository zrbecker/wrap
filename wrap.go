package wrap

type wrapError struct {
	err error
}

type Result[T any] struct {
	value T
	err   error
}

func OK[T any](value T) Result[T] {
	return Result[T]{value: value}
}

func Error[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) UnwrapOrError() (T, error) {
	return r.value, r.err
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(wrapError{err: r.err})
	}
	return r.value
}

func UpwrapHandler[T any](result *Result[T]) {
	if r := recover(); r != nil {
		wrapErr, ok := r.(wrapError)
		if !ok {
			panic(r)
		}
		*result = Result[T]{err: wrapErr.err}
	}
}
