package p_interface

type Scalable[T any] interface {
	Scale(value float64) T
}
