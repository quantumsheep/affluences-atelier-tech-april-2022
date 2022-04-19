package main

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number | ~string](a T, b T) T {
	return a + b
}

func Sub[T Number](a T, b T) T {
	return a - b
}

func Mul[T Number](a T, b T) T {
	return a * b
}

func Div[T Number](a T, b T) T {
	return a / b
}

func Mod[T constraints.Integer](a T, b T) T {
	return a % b
}
