package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return 0.0
	}
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//return min item in arry
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0.0
	}
	result := xs[0]
	for _, item := range xs {
		if item < result {
			result = item
		}
	}
	return result
}

//return max item in arry
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0.0
	}
	result := 0.0
	for _, item := range xs {
		if item > result {
			result = item
		}
	}
	return result
}
