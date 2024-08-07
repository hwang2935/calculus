package bjw

import cc "calculus"

func OneHot(size int, idx int) cc.Point {
	res := make([]float64, size)
	res[idx] = 1
	return cc.Point(res)
}

func jacobian(f cc.Function, p cc.Point) cc.Matrix {
	h := 1e-10
	dim := f.Dim()
	res := make([]cc.Point, dim)

	f0 := f.Map(p)
	for i := 0; i < dim; i++ {
		hi := OneHot(dim, i).Scale(h)
		fi := f.Map(p.Add(hi))
		row := make(cc.Point, dim)
		for j := 0; j < dim; j++ {
			row[j] = (fi[j] - f0[j]) / h
		}
		res[i] = row
	}

	return res
}
