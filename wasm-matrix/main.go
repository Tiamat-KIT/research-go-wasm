package main

import (
	// "time"
	ut "utakata_matrix/types"
	ur "utakata_matrix/routine"
)

type MatrixProps = ut.MatrixProps[int]
type Matrix = ut.Matrix[int]
type MatrixOneLineFuncType = ut.MatrixOneLineFuncType

func main(){}

//export add
func add(p MatrixProps) Matrix{

	// 行列の計算結果
	var results = make([]int, len(p.A))
	
	// 計算関数の複製を行う
	functions := ur.CreateRoutineBaseFunctions(len(p.A))
	
	// ルーチン処理に投げて実際の計算を開始する
	ur.GoRoutine(functions)

	var result int
	for result := range functions {
		results = append(results, result)
	}
	println("results:",result)	
	return result
}
