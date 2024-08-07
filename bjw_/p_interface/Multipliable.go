package p_interface

type Multipliable[T1 any, T2 any] interface {
	Multiply(other T1) T2
}
