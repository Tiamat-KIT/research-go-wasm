package routine

import (
	"sync"
	um "utakata_matrix/base_func"
	ut "utakata_matrix/types"
)

// ルーチンとして扱う複数の処理を生成
func CreateRoutineFunctions[T int | float64](routine_num int) (functions []func([]T,[]T) []T){
	for i := 0; i < routine_num; i++ {
		functions = append(functions,um.Matrix_add[T])
	}
	return functions
}

func GoRoutine[T int|float64](funcList []func([]T,[]T) []T,matrix ut.MatrixProps[T]) ut.Matrix[T] {
	println("started.")
	// 空の行列を作成する
	var results = make(ut.Matrix[T], len(matrix.A))

	// 実行待機している関数のグループ
	var waitGroup sync.WaitGroup

	for idx,function := range funcList{
		waitGroup.Add(1)
		go func(f func([]T,[]T) []T){
			defer waitGroup.Done()
			f(matrix.A[idx],matrix.B[idx])
		}(function)
	}

	waitGroup.Wait()
	println("all finished")
	return results
}