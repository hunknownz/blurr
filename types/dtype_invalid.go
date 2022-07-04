package types

import "log"

type DTInvalid struct {
	nanValue Element
}

func (e *DTInvalid) DType() DType {
	return DInvalid
}

func (e *DTInvalid) Int() (int, bool) {
	return 0, false
}

func (e *DTInvalid) MustInt(args ...int) int {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return 0
}
