package sih

import (
	"calculus"
	"math"
)

const (
	eps = 1e-6
	lim = 1e-4
)

type SingleVariable struct {
	calculus.Domain
	Value func(float64) float64
}

var _ calculus.Function = SingleVariable{}

func (s SingleVariable) Map(p calculus.Point) calculus.Point {
	x := p[0]
	return calculus.Point{s.Value(x)}
}

func (s SingleVariable) Eval(p calculus.Point) float64 {
	return s.Value(p[0])
}

func (s SingleVariable) Diff(x float64) float64 {
	left := s.Value(x - eps)
	right := s.Value(x + eps)
	y := s.Value(x)
	leftp := (y - left) / eps
	rightp := (right - y) / eps
	if math.Abs(leftp-rightp) > lim {
		return math.NaN()
	}
	return (rightp + leftp) / 2
}
