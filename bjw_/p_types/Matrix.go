package p_type

import (
	"calculus/bjw_/p_interface"
	"fmt"
)

type if_matrix interface {
	Object
	p_interface.Scalable[Matrix]
	p_interface.Addable[Matrix]
	p_interface.Multipliable[Matrix, Matrix]
}

type Matrix struct {
	if_matrix
	elements [][]float64
}

func NewMatrix(width int, height int, elements ...float64) Matrix {
	m := Matrix{
		elements: make([][]float64, height),
	}
	for i := range m.elements {
		m.elements[i] = make([]float64, width)
	}

	if len(elements) > 0 {
		k := 0
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if k < len(elements) {
					m.elements[i][j] = elements[k]
					k++
				}
			}
		}
	}

	return m
}

func (m Matrix) Print() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%f ", m.elements[i][j])
		}
		fmt.Println()
	}
}

func TestMatrix() {
	// 예시: 2x2 행렬 생성
	matrix := NewMatrix(2, 2, 1.0, 2.0, 3.0, 4.0)

	// 행렬 출력
	matrix.Print()
}
