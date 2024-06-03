package routine

import (
	"sync"
	um "utakata_matrix/base_func"
	ut "utakata_matrix/types"
)

func CreateRoutineBaseFunctions(routine_num int) (functions []ut.MatrixOneLineFuncType){
	for i := 0; i < routine_num; i++ {
		functions = append(functions,CreateFunctionBase())
	}
	return functions
}



func CreateFunctionBase() ut.MatrixOneLineFuncType {
	return um.Matrix_add[int]
}

func GoRoutine(funcList []ut.MatrixOneLineFuncType) error {
	println("started.")

	// 実行待機している関数のグループ
	var waitGroup sync.WaitGroup

	for _,function := range funcList{
		waitGroup.Add(1)
		go func(f ut.MatrixOneLineFuncType){
			defer waitGroup.Done()
			f()
		}(function)
	}

	waitGroup.Wait()
	println("all finished")
	return nil
}