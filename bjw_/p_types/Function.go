package p_type

import (
	"fmt"
	"math"
)

type if_function interface {
	Object
	Computation(p Vector) Vector
	inputDim() int
	outputDim() int
}

type Function struct {
	if_function
	_function  func(p Vector) Vector
	_inputDim  int
	_outputDim int
}

func (f *Function) Computation(v Vector) Vector {
	if v.Dimension() != f._inputDim {
		panic("함수의 입력 사이즈와 벡터가 일치하지 않습니다.")
	}
	return f._function(v)
}

func NewFunction(function func(p Vector) Vector, inputSize int, outputSize int) Function {
	return Function{
		_function:  function,
		_inputDim:  inputSize,
		_outputDim: outputSize,
	}
}

func (f Function) Jacobian(v Vector) Matrix {
	h := 1e-10
	dim := v.Dimension()
	res := make([]float64, dim*dim)

	f0 := f.Computation(v)
	for i := 0; i < dim; i++ {
		hi := OneHot(dim, i).Scale(h)
		fi := f.Computation(v.Add(hi))
		for j := 0; j < dim; j++ {
			res[i*dim+j] = (fi.elements[j] - f0.elements[j]) / h
		}
	}

	return NewMatrix(dim, dim, res...)
}

func TestFunction() {
	v := NewVector(1, 2, 3)

	f := NewFunction(func(p Vector) Vector {
		sum := 0.
		for _, value := range p.elements {
			sum += value
		}
		return NewVector(sum)
	}, 3, 1)

	result := f.Computation(v)
	fmt.Printf("Result: %+v\n", result)
}

func TestJacobian() {
	x := NewVector(1, 2)

	f := NewFunction(func(v Vector) Vector {
		x1 := v.elements[0]
		x2 := v.elements[1]

		y1 := x1*x1 + x2
		y2 := math.Sin(x1) + math.Cos(x2)
		return NewVector(y1, y2)
	}, 2, 2)

	f.Jacobian(x).Print()
}
