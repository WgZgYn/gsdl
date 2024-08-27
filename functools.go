package main

import (
	"errors"
)

func exec(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}

func runs(fs ...func() error) {
	for _, f := range fs {
		exec(f)
	}
}

func unwrap[T any](f func() (T, error)) T {
	if v, err := f(); err != nil {
		panic(err)
	} else {
		return v
	}
}

func expect[T any](f func() (T, error), s string) T {
	if v, err := f(); err != nil {
		panic(errors.Join(errors.New(s), err))
	} else {
		return v
	}
}

func filter[T any](f func() (T, error), e func(error)) T {
	v, err := f()
	e(err)
	return v
}

func or[T any](f func() (T, error), t T) T {
	if v, err := f(); err != nil {
		return t
	} else {
		return v
	}
}

func input2[A, B, C any](f func(A, B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}

func input1[A, B any](f func(A) B, a A) func() B {
	return func() B {
		return f(a)
	}
}

func wrap[A, B, C any](f func(A) (B, C), a A) func() (B, C) {
	return func() (B, C) {
		return f(a)
	}
}

type Option[T any] struct {
	val *T
}

func (op *Option[T]) unwrap() T {
	if op.val != nil {
		return *op.val
	}
	panic(errors.New("unwrap on nil"))
}

type Pair[F, S any] struct {
	first  F
	second S
}

func ignore(f func() error) {
	_ = f()
}
