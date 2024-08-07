package kdy //이변수 함수에 대한 적분으로 변경.

import (
	"calculus"
	"math"
)

const N = 1000 //적분 영역 세분화

func ValidNumber(f float64) bool {
	return !math.IsInf(f, 0) && !math.IsNaN(f)
}

func Countinuity(f calculus.RealValued, p calculus.Point) bool {
	for i := 0; i < len(p); i++ {
		q := make(calculus.Point, len(p))
		q[i] = p[i] + Eps
		plus := f.Eval(q)
		q[i] = p[i] - Eps
		minus := f.Eval(p)
		if !ValidNumber(plus) || !ValidNumber(minus) {
			return false
		}
		if math.Abs(plus-minus) > Eps {
			return false
		}
	}
	return true
}

func Integral(f calculus.RealValued, bounds []float64) float64 {
	if f.Dim() != 2 {
		panic("Integral: f must be a two-variable function")
	}
	dx := (bounds[1] - bounds[0]) / N
	dy := (bounds[3] - bounds[2]) / N
	var sum float64
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := bounds[0] + float64(i)*dx
			y := bounds[2] + float64(j)*dy
			p := calculus.Point{x, y}
			sum += f.Eval(p) * dx * dy
		}
	}
	return sum
}
