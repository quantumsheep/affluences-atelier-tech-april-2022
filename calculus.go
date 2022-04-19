package main

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Add[T constraints.Integer | ~string](a T, b T) T {
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

func Mod[T constraints.Integer](a T, b T) T {
	return a % b
}

func Pow[T constraints.Integer](a T, b T) float64 {
	return math.Pow(float64(a), float64(b))
}
