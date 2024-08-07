package calculus

import "math"

type Point []float64

func (p Point) Add(q Point) Point {
	dim := len(p)
	r := make([]float64, dim)
	for i := 0; i < dim; i++ {
		r[i] = p[i] + q[i]
	}
	return r
}

func (p Point) AdditiveInverse() Point {
	dim := len(p)
	r := make([]float64, dim)
	for i := 0; i < dim; i++ {
		r[i] = -p[i]
	}
	return r
}

func (p Point) Scale(k float64) Point {
	q := make(Point, len(p))
	for i := 0; i < len(p); i++ {
		q[i] = k * p[i]
	}
	return q
}

func (p Point) Inner(q Point) float64 {
	dim := len(p)
	var inner float64
	for i := 0; i < dim; i++ {
		inner += p[i] * q[i]
	}
	return inner
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt(p.Inner(q))
}
