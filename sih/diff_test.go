package sih_test

import (
	"calculus"
	"calculus/sih"
	"testing"
)

type Interval [2]float64

var _ calculus.Domain = Interval{}

func (i Interval) Contains(p calculus.Point) bool {
	if len(p) == 1 {
		return i[0] <= p[0] && p[0] <= i[1]
	}
	return false
}

func (i Interval) Bounds() []float64 {
	return []float64{i[0], i[1]}
}

func (i Interval) Dim() int {
	return 1
}

func TestSingleVariable_Diff(t *testing.T) {
	f := sih.SingleVariable{
		Domain: Interval{0, 1},
		Value: func(x float64) float64 {
			return x * x
		},
	}
	t.Log(f.Diff(1.))
}
