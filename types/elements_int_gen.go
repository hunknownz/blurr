// Code generated by "go run gen_elements.go"; DO NOT EDIT.
package types

import (
	"errors"
	"fmt"
)

func (es *ElementsInt) Size() int {
	return len(*es)
}

func (es *ElementsInt) Clone() Elements {
	nArray := make(ElementsInt, es.Size())
	copy(nArray, *es)
	return &nArray
}

func (es *ElementsInt) DType() DType {
	return DInt
}

func (es *ElementsInt) IGet(i int) (e Element, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	e = (*es)[i]
	return
}

func (es *ElementsInt) ISet(i int, e Element) (isNaN bool, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	var v int
	v, isNaN = e.Int()
	if isNaN {
		return
	}
	(*es)[i] = DTInt(v)
	return
}

func (es *ElementsInt8) Size() int {
	return len(*es)
}

func (es *ElementsInt8) Clone() Elements {
	nArray := make(ElementsInt8, es.Size())
	copy(nArray, *es)
	return &nArray
}

func (es *ElementsInt8) DType() DType {
	return DInt8
}

func (es *ElementsInt8) IGet(i int) (e Element, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	e = (*es)[i]
	return
}

func (es *ElementsInt8) ISet(i int, e Element) (isNaN bool, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	var v int8
	v, isNaN = e.Int8()
	if isNaN {
		return
	}
	(*es)[i] = DTInt8(v)
	return
}

func (es *ElementsInt16) Size() int {
	return len(*es)
}

func (es *ElementsInt16) Clone() Elements {
	nArray := make(ElementsInt16, es.Size())
	copy(nArray, *es)
	return &nArray
}

func (es *ElementsInt16) DType() DType {
	return DInt16
}

func (es *ElementsInt16) IGet(i int) (e Element, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	e = (*es)[i]
	return
}

func (es *ElementsInt16) ISet(i int, e Element) (isNaN bool, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	var v int16
	v, isNaN = e.Int16()
	if isNaN {
		return
	}
	(*es)[i] = DTInt16(v)
	return
}

func (es *ElementsInt32) Size() int {
	return len(*es)
}

func (es *ElementsInt32) Clone() Elements {
	nArray := make(ElementsInt32, es.Size())
	copy(nArray, *es)
	return &nArray
}

func (es *ElementsInt32) DType() DType {
	return DInt32
}

func (es *ElementsInt32) IGet(i int) (e Element, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	e = (*es)[i]
	return
}

func (es *ElementsInt32) ISet(i int, e Element) (isNaN bool, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	var v int32
	v, isNaN = e.Int32()
	if isNaN {
		return
	}
	(*es)[i] = DTInt32(v)
	return
}

func (es *ElementsInt64) Size() int {
	return len(*es)
}

func (es *ElementsInt64) Clone() Elements {
	nArray := make(ElementsInt64, es.Size())
	copy(nArray, *es)
	return &nArray
}

func (es *ElementsInt64) DType() DType {
	return DInt64
}

func (es *ElementsInt64) IGet(i int) (e Element, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	e = (*es)[i]
	return
}

func (es *ElementsInt64) ISet(i int, e Element) (isNaN bool, err error) {
	if i > es.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", es.Size()))
		return
	}
	var v int64
	v, isNaN = e.Int64()
	if isNaN {
		return
	}
	(*es)[i] = DTInt64(v)
	return
}
