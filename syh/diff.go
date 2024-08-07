package syh

import (
	"calculus"
	"errors"
)

const eps = 1e-6

func Partial(f calculus.RealValued, p calculus.Point, d int) (float64, error) {
	if d < 0 || d >= len(p) {
		return 0, errors.New("direction out of range")
	}
	p[d] += eps
	fxPlusEps := f.Eval(p)
	p[d] -= 2 * eps
	fxMinusEps := f.Eval(p)
	p[d] += eps
	df := (fxPlusEps - fxMinusEps) / (2 * eps)
	return df, nil
}
