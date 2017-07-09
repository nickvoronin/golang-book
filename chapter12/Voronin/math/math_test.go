package math

import "testing"

type testset struct {
	values []float64
	average float64
	max float64
	min float64
}

var tests = []testset{
	{ []float64{1,2}, 1.5, 2, 1 },
	{ []float64{1,1,1,1,1,1}, 1, 1, 1 },
	{ []float64{-1,1}, 0, 1, -1 },
	{ []float64{}, 0, 0, 0 },
}

func TestAverage(t *testing.T) {
	for _, set := range tests {
		v := Average(set.values)
		if v != set.average {
			t.Error(
				"For", set.values,
				"expected", set.average,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, set := range tests {
		max := Max(set.values)
		if max != set.max {
			t.Error(
				"For", set.values,
				"expected", set.max,
				"got", max,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, set := range tests {
		min := Min(set.values)
		if min != set.min {
			t.Error(
				"For", set.values,
				"expected", set.min,
				"got", min,
			)
		}
	}
}