package kdy_test

import (
	"testing"
	"calculus/kdy"
)

func TestIntegral(t *testing.T) {
	f := kdy.TwoVariable(func(x, y float64) float64 {
		return x*x + y*y
	})
	t.Log(kdy.Integral(f, []float64{0, 1, 0, 1}))
}