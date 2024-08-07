package p_type

import (
	"calculus/bjw_/p_interface"
	"fmt"
	"math"
)

type if_vector interface {
	if_point
	p_interface.Addable[Vector]
	p_interface.Multipliable[Vector, Matrix]
}

type Vector struct {
	if_vector
	elements []float64
}

func NewVector(elements ...float64) Vector {
	var v Vector
	v.elements = append(v.elements, elements...)
	return v
}

func (v Vector) Dimension() int {
	return len(v.elements)
}

func (v Vector) Add(other Vector) Vector {
	if v.Dimension() != other.Dimension() {
		panic("벡터의 차원이 다릅니다")
	}

	result := make([]float64, v.Dimension())
	for i := range v.elements {
		result[i] = v.elements[i] + other.elements[i]
	}
	return NewVector(result...)
}

func (v Vector) Scale(scalar float64) Vector {
	result := make([]float64, v.Dimension())
	for i := range v.elements {
		result[i] = v.elements[i] * scalar
	}
	return NewVector(result...)
}

func (v Vector) Magnitude() float64 {
	sum := 0.0
	for _, component := range v.elements {
		sum += component * component
	}
	return math.Sqrt(sum)
}

func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return NewVector(make([]float64, v.Dimension())...)
	}
	result := make([]float64, v.Dimension())
	for i := range v.elements {
		result[i] = v.elements[i] / magnitude
	}
	return NewVector(result...)
}

func OneHot(size int, index int) Vector {
	elements := make([]float64, size)
	if index >= 0 && index < size {
		elements[index] = 1
	} else {
		panic("index cannot be bigger than size or equals")
	}
	return NewVector(elements...)
}

func (v Vector) DotProduct(other Vector) (float64, error) {
	if v.Dimension() != other.Dimension() {
		return 0, fmt.Errorf("벡터의 차원이 다릅니다")
	}
	sum := 0.0
	for i := range v.elements {
		sum += v.elements[i] * other.elements[i]
	}
	return sum, nil
}

func TestVector() {
	v1 := NewVector(3, 4, 5)
	v2 := NewVector(1, 2, 3)

	// 벡터 덧셈
	sum := v1.Add(v2)
	fmt.Printf("Sum: %+v\n", sum)

	// 스칼라 곱셈
	scaled := v1.Scale(2)
	fmt.Printf("Scaled: %+v\n", scaled)

	// 벡터 크기
	magnitude := v1.Magnitude()
	fmt.Printf("Magnitude: %f\n", magnitude)

	// 벡터 정규화
	normalized := v1.Normalize()
	fmt.Printf("Normalized: %+v\n", normalized)

	// 내적
	dotProduct, err := v1.DotProduct(v2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Dot Product: %f\n", dotProduct)
	}
}

func TestVector1() {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(4, 5, 6)
	result := v1.Multiply(v2)

	fmt.Printf("Result: %+v\n", result)
}
