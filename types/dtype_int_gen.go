// Code generated by "go run gen_dtype.go"; DO NOT EDIT.
package types

import "unsafe"

func (dt DTInt) DType() DType {
	return DInt
}

func (dt DTInt) ItemSize() int {
	return int(unsafe.Sizeof(dt))
}

func (dt DTInt8) DType() DType {
	return DInt8
}

func (dt DTInt8) ItemSize() int {
	return int(unsafe.Sizeof(dt))
}

func (dt DTInt16) DType() DType {
	return DInt16
}

func (dt DTInt16) ItemSize() int {
	return int(unsafe.Sizeof(dt))
}

func (dt DTInt32) DType() DType {
	return DInt32
}

func (dt DTInt32) ItemSize() int {
	return int(unsafe.Sizeof(dt))
}

func (dt DTInt64) DType() DType {
	return DInt64
}

func (dt DTInt64) ItemSize() int {
	return int(unsafe.Sizeof(dt))
}
