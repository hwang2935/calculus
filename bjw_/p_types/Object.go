package p_type

import "calculus/bjw_/p_interface"

type if_object interface {
	p_interface.Equalable
}

type Object interface {
	if_object
}
