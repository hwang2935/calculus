package sgw

import "calculus"

func Cross(v1, v2 calculus.Point) calculus.Point {
	if len(v1) != 3 || len(v2) != 3 {
		panic("Cross: Dimension must be 3.")
	}
	return calculus.Point{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}
}
