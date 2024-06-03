package base_func

func Matrix_add[T int | float64](a []T,b []T) []T {
	result := make([]T,len(a))
	for i := 0; i < len(a); i++{
		result[i] = a[i] + b[i]
	}
	return result
}

func Matrix_sub[T int | float64](a []T,b []T) []T {
	result := make([]T,len(a))
	for i := 0; i < len(a); i++{
		result[i] = a[i] - b[i]
	}
	return result
}