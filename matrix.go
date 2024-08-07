package calculus

import "math"

type Matrix []Point

func (m Matrix) Add(n Matrix) Matrix {
	if len(m) != len(n) || len(m[0]) != len(n[0]) {
		panic("Matrix.Add: Dimensions mismatch")
	}
	r, c := len(m), len(m[0])
	l := make(Matrix, r)
	for i := 0; i < r; i++ {
		l[i] = make(Point, c)
		for j := 0; j < c; j++ {
			l[i] = m[i].Add(n[i])
		}
	}
	return l
}

func (m Matrix) AdditiveInverse() Matrix {
	r, c := len(m), len(m[0])
	l := make(Matrix, r)
	for i := 0; i < r; i++ {
		l[i] = make(Point, c)
		for j := 0; j < c; j++ {
			l[i] = m[i].AdditiveInverse()
		}
	}
	return l
}

func (m Matrix) Scale(s float64) Matrix {
	r, c := len(m), len(m[0])
	l := make(Matrix, r)
	for i := 0; i < r; i++ {
		l[i] = make(Point, c)
		for j := 0; j < c; j++ {
			l[i] = m[i].Scale(s)
		}
	}
	return l
}

func (m Matrix) Transpose() Matrix {
	r, c := len(m[0]), len(m)
	l := make(Matrix, r)
	for i := 0; i < r; i++ {
		l[i] = make(Point, c)
		for j := 0; j < c; j++ {
			l[i][j] = m[j][i]
		}
	}
	return l
}

func (m Matrix) Mul(n Matrix) Matrix {
	if len(m[0]) != len(n) {
		panic("Matrix.Mul: Dimensions mismatch")
	}
	r, c := len(m), len(n[0])
	l := make(Matrix, r)
	for i := 0; i < r; i++ {
		l[i] = make(Point, c)
		for j := 0; j < c; j++ {
			l[i][j] = m[i].Inner(n.Transpose()[j])
		}
	}
	return l
}

func (m Matrix) Transform(p Point) Point {
	q := make(Point, len(p))
	for i := 0; i < len(p); i++ {
		q[i] = m[i].Inner(p)
	}
	return q
}

func Rotation3D(n int, theta float64) Matrix {
	m := make(Matrix, 3)
	for i := 0; i < 3; i++ {
		m[i] = make(Point, 3)
	}
	m[n][n] = 1
	i, j := (n+1)%3, (n+2)%3
	m[i][i] = math.Cos(theta)
	m[i][j] = -math.Sin(theta)
	m[j][i] = math.Sin(theta)
	m[j][j] = math.Cos(theta)
	return m
}
