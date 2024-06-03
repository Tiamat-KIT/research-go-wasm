package main

import (
	// "time"
	ut "utakata_matrix/types"
	ur "utakata_matrix/routine"
	"fmt"
)

type MatrixProps = ut.MatrixProps[int]
type Matrix = ut.Matrix[int]

func main(){
	var result = add[int](MatrixProps{
		A: [][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		},
		B: [][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
		},
	})
	fmt.Println(result)
}

//export add
func add[T int | float64](p ut.MatrixProps[T]) ut.Matrix[T]{
	// ルーチン処理に投げて実際の計算を開始する
	return ur.GoRoutine[T](ur.CreateRoutineFunctions[T](len(p.A)),p)
}
