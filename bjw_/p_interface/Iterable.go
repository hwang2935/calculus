package p_interface

type Iterable interface {
	Iter() <-chan interface{}
}
