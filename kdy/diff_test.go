package kdy_test

import (
	"calculus/kdy"
	"testing"
)

func TestTwoVariable_Partial(t *testing.T) {
	f := kdy.TwoVariable(func(x, y float64) float64 {
		return x*x + y*y
	})
	t.Log(f.Partial(1., 1., 0))
}

func TestPartial(t *testing.T) {
	f := kdy.TwoVariable(func(x, y float64) float64 {
		return x*x + y*y
	})
	t.Log(kdy.Partial(f, []float64{1., 1.}, 0))
}
