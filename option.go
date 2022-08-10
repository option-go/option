package option

import "errors"

var errorOptionIsNone error = errors.New("Option has null value")

type Option[T any] struct {
	Some T
	None bool
}

func Some[T any](some T) Option[T] {
	return Option[T]{
		Some: some,
		None: false,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		None: true,
	}
}

func (opt *Option[T]) IsSome() bool {
	return !opt.None
}

func (opt *Option[T]) IsNone() bool {
	return opt.None
}

func (opt *Option[T]) Unwrap() T {
	if opt.None == true {
		panic(errorOptionIsNone)
	}
	return opt.Some
}
