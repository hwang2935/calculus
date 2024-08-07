package lhj

import (
	"calculus"
	"reflect"
)

type GeneralSet struct {
	ElementType reflect.Type
}

func (N GeneralSet) Contains(x any) bool {
	return reflect.TypeOf(x) == N.ElementType
}

type Function struct {
	Domain   GeneralSet
	Codomain GeneralSet
	Relation func(x calculus.Point) calculus.Matrix
}

func (f Function) Computation(v calculus.Point) calculus.Matrix {
	ok := true
	if !ok {
		return calculus.Matrix{}
	}
	return f.Relation(v)
}

const epsilon = 1e-6
const bigNum = 1000000

func Diff(f Function) func(calculus.Point) calculus.Matrix {
	return func(x calculus.Point) calculus.Matrix {
		if !IsSafe(x, f.Domain) {
			panic("Diff: x is not in the domain of f")
		}
		domainDim := len(x)
		totalderiv := make([]calculus.Point, domainDim)
		for i := range x {
			infinitesimal := make([]float64, domainDim)
			infinitesimal[i] = epsilon
			epsilonCoord := calculus.Point(infinitesimal)
			slope := (f.Computation(x.Add(epsilonCoord)).Add(f.Computation(x).AdditiveInverse())).Scale(bigNum)
			totalderiv[i] = calculus.Point(slope[0])
		}

		return calculus.Matrix(totalderiv)
	}
}

func IsSafe(input any, domain GeneralSet) bool {
	return domain.Contains(input)
}

/*

  package lhj

  type Set interface {
	Type(any) bool
  }

  type RealSet struct {}

  var _ Set = RealSet{}

  func (s RealSet) Contain(x any) bool {
	if y, ok := x.(float64); ok {
	  return true
	}
	return false
  }

  var _ calculus.Domain = RealSet{}

  func (s RealSet) Contains(p Point) bool {
	if len(p) == 1 {
	  return true
	}
	return false
  }

  func (s RealSet) Bounds() []float64 {
	return []float64{math.Inf(-1), math.Inf(1)}
  }

  func (s RealSet) Dim() int {
	return 1
  }

  type Interval struct {
	RealSet
	Ends [2]float64
  }

  func (i Interval) Bounds() []float64 {
	return []float64{i.Ends[0], i.Ends[1]}
  }


  type Function struct {
	Domain
	Assign func(calculus.Point) calculus.Point
  }

  var _ calculus.Function = Function{}

  func (f Function) Map(p calculus.Point) calculus.Point
	return f.Assign(p)
  }



  func Differential(F Function) Function {
	var Derivative Function
	Derivative.domain = F.domain
	Derivative.codomain = F.codomain
	Derivative.maprelation = Diffrel(F, F.domain)
	return Derivative
  }

  type Function interface{}

  type PointFunction struct{}

  type Operator struct{}

  func (o Operator) Partial(i int) Function {
	p := Operator{}

	return Function(p)
  }

  type Differentiable interface {
	Partial(int) Function
  }

  type DifferentiableFunction interface {
	Function
	Differentiable
  }

  type Matrix struct{}

  var _ Function = Matrix{}

  func Diff(f DifferentiableFunction) Function {
	dim := f.Dim()
	mat := make(Matrix, dim)
	for i := 0 ; i < dim ; i++ {
	  mat[i] = Gradient(f)
	}
	return mat
  }

  func Gradient(f DifferentiableFunction) Matrix {
	dim := f.Dim()
	mat := make(Matrix, dim)
	for i := 0 ; i < dim;i++ {
	  mat[i] = f.Partial(i)
	}
	return mat
  }

*/
