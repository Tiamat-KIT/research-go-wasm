package main

func main(){
	println("add two numbers",add(2,5))
}

// export add
func add(a int,b int) int{
	return a+b
}

// export multiply
func multiply(a int,b int) int{
	return a*b
}

// export divide
func divide(a int,b int) int{
	return a/b
}

// export subtract
func subtract(a int,b int) int{
	return a-b
}