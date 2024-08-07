package syh

import (
	"calculus"
)

func Integral(f calculus.RealValued, n int, a, b float64) float64 {
	var sum float64

	if n <= 0 {
		panic("the number of subdivisions must be greater than zero")
	}

	dx := (b - a) / float64(n)

	for i := 0; i < n; i++ {
		x := a + float64(i)*dx + dx/2
		sum += f.Eval(calculus.Point{x}) * dx
	}

	return sum
}
