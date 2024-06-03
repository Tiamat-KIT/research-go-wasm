package main

import (
	// "time"
	ut "utakata_matrix/types"
	ur "utakata_matrix/routine"
)

type MatrixProps = ut.MatrixProps[int]
type Matrix = ut.Matrix[int]

func main(){}

//export add
func add(p MatrixProps) Matrix{
	
	// ルーチン処理に投げて実際の計算を開始する
	var results = ur.GoRoutine[int](ur.CreateRoutineFunctions[int](len(p.A)),p)

	return results
}
