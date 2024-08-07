package lhj_test

import (
	"calculus"
	"calculus/lhj"
	"math"
	"reflect"
	"testing"
)

var Real lhj.GeneralSet = lhj.GeneralSet{
	ElementType: reflect.TypeOf(calculus.Point{0.}),
}

func TestDiff(t *testing.T) {
	f1 := lhj.Function{
		Domain:   Real,
		Codomain: Real,
		Relation: func(p calculus.Point) calculus.Matrix {
			return calculus.Matrix{{math.Cos(p[0])}}
		},
	}
	f2 := lhj.Function{
		Domain:   Real,
		Codomain: Real,
		Relation: func(p calculus.Point) calculus.Matrix {
			x, y := p[0], p[1]
			return calculus.Matrix{{x*x + y*y}}
		},
	}
	t.Log(lhj.Diff(f1)(calculus.Point{1.}))
	t.Log(lhj.Diff(f2)(calculus.Point{1., 1.}))
}
