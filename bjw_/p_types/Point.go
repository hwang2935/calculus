package p_type

import (
	"calculus/bjw_/p_interface"
	"math"
)

type if_point interface {
	if_object
	p_interface.Containable[float64]
	p_interface.Iterable
}

type Point struct {
	if_point
	elements []float64
}

func NewPoint(elements ...float64) Point {
	var p Point
	p.elements = []float64{}
	p.elements = append(p.elements, elements...)
	return p
}

func (p Point) Dimension() int {
	return len(p.elements)
}

func (p Point) Distance(other Point) float64 {
	if p.Dimension() != other.Dimension() {
		panic("Point의 Dimension이 다릅니다.")
	}

	sum := 0.0
	for i := range p.elements {
		diff := p.elements[i] - other.elements[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func (p Point) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for _, value := range p.elements {
			ch <- value
		}
		close(ch)
	}()
	return ch
}
