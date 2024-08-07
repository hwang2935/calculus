package p_type

import (
	"calculus/bjw_/p_interface"
	"reflect"
)

type if_set interface {
	p_interface.Containable[any]
	p_interface.Equalable
}

type Set struct {
	if_set
}

func (s Set) Equals(other any) bool {
	if reflect.TypeOf(other) != reflect.TypeOf(Set{}) {
		return false
	} else {
		return true
	}
}
