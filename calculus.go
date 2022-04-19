package main

import "golang.org/x/exp/constraints"

func Add[T constraints.Ordered](a T, b T) T {
	return a + b
}

func Sub[T constraints.Integer](a T, b T) T {
	return a * b
}

func Mul[T constraints.Integer](a T, b T) T {
	return a * b
}

func Div[T constraints.Integer](a T, b T) T {
	return a / b
}
