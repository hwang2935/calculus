package sih_test

import (
	"calculus"
	"calculus/sih"
	"testing"
)

type Rectangular []Interval

var _ calculus.Domain = Rectangular{}

func (r Rectangular) Contains(p calculus.Point) bool {
	if len(p) == len(r) {
		for i, x := range p {
			if !r[i].Contains(calculus.Point{x}) {
				return false
			}
		}
		return true
	}
	return false
}

func (r Rectangular) Bounds() []float64 {
	var bounds []float64
	for _, interval := range r {
		bounds = append(bounds, interval[0], interval[1])
	}
	return bounds
}

func (r Rectangular) Dim() int {
	return len(r)
}

func TestGraph2D(t *testing.T) {
	f := sih.SingleVariable{
		Domain: Rectangular{
			{-1, 1},
			{-1, 1},
		},
		Value: func(x float64) float64 {
			return x * x
		},
	}
	calculus.Graph2D(f, 600, 600, "graph2d.png")
}
