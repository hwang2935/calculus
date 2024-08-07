package bjw

import cc "calculus"

func generateGrid(current []float64, bounds []float64, dim int, step float64) [][]float64 {
	if dim >= len(bounds)/2 {
		return [][]float64{append([]float64{}, current...)}
	}

	var points [][]float64
	min := bounds[dim*2]
	max := bounds[dim*2+1]

	for value := min; value <= max+step/2; value += step {
		newPoint := append(current, value)
		points = append(points, generateGrid(newPoint, bounds, dim+1)...)
	}

	return points
}

func get_subdivision_points(bounds []float64, step float64) [][]float64 {
	if len(bounds)%2 != 0 {
		panic("Bounds array must contain even number of elements")
	}
	return generateGrid([]float64{}, bounds, 0, step)
}

func get_box_size(step_size float64, dim int) float64 {
	res := 1.
	for i := 0; i < dim; i++ {
		res *= step_size
	}
	return res
}

func integral(f cc.RealValued, step float64) float64 {
	res := 0.
	dim := f.Dim()
	bound := f.Bounds()

	box := get_box_size(step, dim)

	grid_points = get_subdivision_points(boun, step)
	for i, point := range grid_points{
		res += f.Map(point) * box
	}

	return res
}