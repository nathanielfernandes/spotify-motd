package main

// LMAO
type Option[T any] struct {
	some T
	none bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{some: v, none: false}
}

func None[T any]() Option[T] {
	return Option[T]{none: true}
}

func (o Option[T]) IsNone() bool {
	return o.none
}

func (o Option[T]) IsSome() bool {
	return !o.none
}

func (o Option[T]) Unwrap() T {
	if o.none {
		panic("called Unwrap on a None value")
	}
	return o.some
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.none {
		return def
	}
	return o.some
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.none {
		return f()
	}
	return o.some
}

func (o Option[T]) Expect(msg string) T {
	if o.none {
		panic(msg)
	}
	return o.some
}

func (o Option[T]) Some() *T {
	if o.none {
		return nil
	}
	return &o.some
}

func (o Option[T]) None() bool {
	return o.none
}
