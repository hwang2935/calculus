package sgw_test

import (
	"calculus"
	"calculus/sgw"
	"testing"
)

func TestCross(t *testing.T) {
	v1 := calculus.Point{1, 2, 3}
	v2 := calculus.Point{4, 5, 6}
	t.Log(sgw.Cross(v1, v2))
}
