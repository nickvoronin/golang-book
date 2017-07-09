package math

// Returns an average value for the given array of floats
func Average(numbers []float64) float64 {
	total := float64(0)
	for _, x := range numbers {
		total += x
	}
	return total / float64(len(numbers))
}

// Returns a maximum value for a given array of floats
func Max(numbers []float64) float64 {
	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	return max
}

// Returns a minimum value for a given array of floats
func Min(numbers []float64) float64 {
	min := numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
	}
	return min
}