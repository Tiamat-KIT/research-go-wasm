package main

func main(){}

//export add
func add(a int,b int) int{
	return a+b
}

//export multiply
func multiply(a int,b int) int{
	return a*b Z
}

//export divide
func divide(a int,b int) int{
	return a/b
}

//export subtract
func subtract(a int,b int) int{
	return a-b
}

//export square
func square(a int) int{
	return a*a
}

//export matrix_add
func matrix_add(a [][]int,b [][]int) [][]int{
	// 数値の配列の配列で構成されていると考えた時、
	// 行列の加算を行う関数
	// 対応する行ごとの要素の加算を行うことは
	// 動的なルーチン作成によって、高速な行列計算を実現できる。

}

