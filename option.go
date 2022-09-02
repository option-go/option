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

func (opt *Option[T]) Unwrap() *T {
	if opt.None == true {
		panic(errorOptionIsNone)
	}
	return &opt.Some
}

// Returns None if the option is None, otherwise returns optb.
// Due to go limitations (method must have no type parameters other than specified for receiver)
// this method can be used only on options of same type.
// For options of different type use function OptAndOpt instead.
func (opt *Option[T]) And(optb *Option[T]) *Option[T] {
	return OptAndOpt(opt, optb)
}

// Returns None if the option opta is None, otherwise returns optb.
func OptAndOpt[T, U any](opta *Option[T], optb *Option[U]) *Option[U] {
	if opta.IsSome() == true {
		return optb
	} else {
		none := None[U]()
		return &none
	}
}
