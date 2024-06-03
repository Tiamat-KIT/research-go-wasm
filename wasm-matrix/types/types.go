package types

type Matrix[T int | float64] [][]T

type MatrixProps[T int | float64] struct {
	A Matrix[T]
	B Matrix[T]
}

type MatrixOneLineFuncType func([]int,[]int) []int