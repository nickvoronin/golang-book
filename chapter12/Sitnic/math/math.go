package math

func Min(x ...float64) float64 {
	minNumber := x[0]
	for _, y := range x {
		if y < minNumber {
			minNumber = y
		}
	}

	return minNumber
}

func Max(x ...float64) float64 {
	maxNumber := x[0]
	for _, y := range x {
		if y > maxNumber {
			maxNumber = y
		}
	}

	return maxNumber
}
