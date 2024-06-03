関数を生成して実行する準備を行う関数
createFunctions(num int) (function []MatrixOneLineFuncType)
- for文で関数を作成して格納する
- それを返す

実際の処理の関数を返す関数
createFunctionBase() MatrixOneLineFuncType
- 実際に処理する関数を定義する
- それを返す

加算する関数
add_func = func(a []int,b []int) []int
- 実際の加算を行う関数
    - まず、空の配列を作成する
    - for文で配列の要素の加算を行う
- それを返す

ルーティンを作成して実行していく関数
goRoutine(funcList []MatrixOneLineFuncType) error