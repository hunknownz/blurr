package types

import "math"

type Element interface {
	DType() DType
	Int() (int, bool)
	MustInt(args ...int) int
	Int8() (int8, bool)
	MustInt8(args ...int8) int8
	Int16() (int16, bool)
	MustInt16(args ...int16) int16
	Int32() (int32, bool)
	MustInt32(args ...int32) int32
	Int64() (int64, bool)
	MustInt64(args ...int64) int64
	Float32() (float32, bool)
	MustFloat32(args ...float32) float32
	Float64() (float64, bool)
	MustFloat64(args ...float64) float64
	ItemSize() int
}

func floatGreater(a, b float64) bool {
	return math.Max(a, b) == a && math.Abs(a-b) > 0.00001
}

func floatLess(a, b float64) bool {
	return math.Max(a, b) == b && math.Abs(a-b) > 0.00001
}
