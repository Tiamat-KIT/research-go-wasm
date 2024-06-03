package types

// 行列の型は整数でも浮動小数点でも良い
type Matrix[T int | float64] [][]T

// 行列の演算をする関数の引数型
type MatrixProps[T int | float64] struct {
	A Matrix[T]
	B Matrix[T]
}

// type MatrixOneLineFuncType[T int | float64] func([]T,[]T) []T