//go:generate go run gen_dtype.go -type DTInt*
package types

import (
	"log"
	"math"
	"unsafe"
)

type DTInt int

func (dt DTInt) Int() (int, bool) {
	return int(dt), true
}

func (dt DTInt) MustInt(args ...int) int {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int(dt)
}

func (dt DTInt) Int8() (int8, bool) {
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return 0, false
	}
	return int8(dt), true
}

func (dt DTInt) MustInt8(args ...int8) int8 {
	var def int8
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return def
	}
	return int8(dt)
}

func (dt DTInt) Int16() (int16, bool) {
	if dt > math.MaxInt16 || dt < math.MinInt16 {
		return 0, false
	}
	return int16(dt), true
}

func (dt DTInt) MustInt16(args ...int16) int16 {
	var def int16
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt16 || dt < math.MinInt16 {
		return def
	}
	return int16(dt)
}

func (dt DTInt) Int32() (int32, bool) {
	if dt.ItemSize() == int(unsafe.Sizeof(0)) {
		return int32(dt), true
	}
	if dt > math.MaxInt32 || dt < math.MinInt32 {
		return 0, false
	}
	return int32(dt), true
}

func (dt DTInt) MustInt32(args ...int32) int32 {
	if dt.ItemSize() == int(unsafe.Sizeof(0)) {
		return int32(dt)
	}
	var def int32
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt32 || dt < math.MinInt32 {
		return def
	}
	return int32(dt)
}

func (dt DTInt) Int64() (int64, bool) {
	return int64(dt), true
}

func (dt DTInt) MustInt64(args ...int64) int64 {
	if len(args) > 1 {
		log.Panicf("MustInt64() received too many arguments %d", len(args))
	}
	return int64(dt)
}

func (dt DTInt) Float32() (float32, bool) {
	return float32(dt), true
}

func (dt DTInt) MustFloat32(args ...float32) float32 {
	if len(args) > 1 {
		log.Panicf("MustFloat32() received too many arguments %d", len(args))
	}
	return float32(dt)
}

func (dt DTInt) Float64() (float64, bool) {
	return float64(dt), true
}

func (dt DTInt) MustFloat64(args ...float64) float64 {
	if len(args) > 1 {
		log.Panicf("MustFloat64() received too many arguments %d", len(args))
	}
	return float64(dt)
}

type DTInt8 int8

func (dt DTInt8) Int() (int, bool) {
	return int(dt), true
}

func (dt DTInt8) MustInt(args ...int) int {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int(dt)
}

func (dt DTInt8) Int8() (int8, bool) {
	return int8(dt), true
}

func (dt DTInt8) MustInt8(args ...int8) int8 {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int8(dt)
}

func (dt DTInt8) Int16() (int16, bool) {
	return int16(dt), true
}

func (dt DTInt8) MustInt16(args ...int16) int16 {
	if len(args) > 1 {
		log.Panicf("MustInt16() received too many arguments %d", len(args))
	}
	return int16(dt)
}

func (dt DTInt8) Int32() (int32, bool) {
	return int32(dt), true
}

func (dt DTInt8) MustInt32(args ...int32) int32 {
	if len(args) > 1 {
		log.Panicf("MustInt32() received too many arguments %d", len(args))
	}
	return int32(dt)
}

func (dt DTInt8) Int64() (int64, bool) {
	return int64(dt), true
}

func (dt DTInt8) MustInt64(args ...int64) int64 {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int64(dt)
}

func (dt DTInt8) Float32() (float32, bool) {
	return float32(dt), true
}

func (dt DTInt8) MustFloat32(args ...float32) float32 {
	if len(args) > 1 {
		log.Panicf("MustFloat64() received too many arguments %d", len(args))
	}
	return float32(dt)
}

func (dt DTInt8) Float64() (float64, bool) {
	return float64(dt), true
}

func (dt DTInt8) MustFloat64(args ...float64) float64 {
	if len(args) > 1 {
		log.Panicf("MustFloat64() received too many arguments %d", len(args))
	}
	return float64(dt)
}

type DTInt16 int16

func (dt DTInt16) Int() (int, bool) {
	return int(dt), true
}

func (dt DTInt16) MustInt(args ...int) int {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int(dt)
}

func (dt DTInt16) Int8() (int8, bool) {
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return 0, false
	}
	return int8(dt), true
}

func (dt DTInt16) MustInt8(args ...int8) int8 {
	var def int8
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return def
	}
	return int8(dt)
}

func (dt DTInt16) Int16() (int16, bool) {
	return int16(dt), true
}

func (dt DTInt16) MustInt16(args ...int16) int16 {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int16(dt)
}

func (dt DTInt16) Int32() (int32, bool) {
	return int32(dt), true
}

func (dt DTInt16) MustInt32(args ...int32) int32 {
	if len(args) > 1 {
		log.Panicf("MustInt32() received too many arguments %d", len(args))
	}
	return int32(dt)
}

func (dt DTInt16) Int64() (int64, bool) {
	return int64(dt), true
}

func (dt DTInt16) MustInt64(args ...int64) int64 {
	if len(args) > 1 {
		log.Panicf("MustInt64() received too many arguments %d", len(args))
	}
	return int64(dt)
}

func (dt DTInt16) Float32() (float32, bool) {
	return float32(dt), true
}

func (dt DTInt16) MustFloat32(args ...float32) float32 {
	if len(args) > 1 {
		log.Panicf("MustFloat32() received too many arguments %d", len(args))
	}
	return float32(dt)
}

func (dt DTInt16) Float64() (float64, bool) {
	return float64(dt), true
}

func (dt DTInt16) MustFloat64(args ...float64) float64 {
	if len(args) > 1 {
		log.Panicf("MustFloat64() received too many arguments %d", len(args))
	}
	return float64(dt)
}

type DTInt32 int32

func (dt DTInt32) Int() (int, bool) {
	return int(dt), true
}

func (dt DTInt32) MustInt(args ...int) int {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int(dt)
}

func (dt DTInt32) Int8() (int8, bool) {
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return 0, false
	}
	return int8(dt), true
}

func (dt DTInt32) MustInt8(args ...int8) int8 {
	var def int8
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return def
	}
	return int8(dt)
}

func (dt DTInt32) Int16() (int16, bool) {
	if dt > math.MaxInt16 || dt < math.MinInt16 {
		return 0, false
	}
	return int16(dt), true
}

func (dt DTInt32) MustInt16(args ...int16) int16 {
	var def int16
	if len(args) > 1 {
		log.Panicf("MustInt16() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt16 || dt < math.MinInt16 {
		return def
	}
	return int16(dt)
}

func (dt DTInt32) Int32() (int32, bool) {
	return int32(dt), true
}

func (dt DTInt32) MustInt32(args ...int32) int32 {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int32(dt)
}

func (dt DTInt32) Int64() (int64, bool) {
	return int64(dt), true
}

func (dt DTInt32) MustInt64(args ...int64) int64 {
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	return int64(dt)
}

func (dt DTInt32) Float32() (float32, bool) {
	return float32(dt), true
}

func (dt DTInt32) MustFloat32(args ...float32) float32 {
	if len(args) > 1 {
		log.Panicf("MustFloat32() received too many arguments %d", len(args))
	}
	return float32(dt)
}

func (dt DTInt32) Float64() (float64, bool) {
	return float64(dt), true
}

func (dt DTInt32) MustFloat64(args ...float64) float64 {
	if len(args) > 1 {
		log.Panicf("MustFloat64() received too many arguments %d", len(args))
	}
	return float64(dt)
}

type DTInt64 int64

func (dt DTInt64) Int() (int, bool) {
	if dt.ItemSize() == int(unsafe.Sizeof(0)) {
		return int(dt), true
	}
	if dt > math.MaxInt || dt < math.MinInt {
		return 0, false
	}
	return int(dt), true
}

func (dt DTInt64) MustInt(args ...int) int {
	if dt.ItemSize() == int(unsafe.Sizeof(0)) {
		return int(dt)
	}
	var def int
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt || dt < math.MinInt {
		return def
	}
	return int(dt)
}

func (dt DTInt64) Int8() (int8, bool) {
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return 0, false
	}
	return int8(dt), true
}

func (dt DTInt64) MustInt8(args ...int8) int8 {
	var def int8
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt8 || dt < math.MinInt8 {
		return def
	}
	return int8(dt)
}

func (dt DTInt64) Int16() (int16, bool) {
	if dt > math.MaxInt16 || dt < math.MinInt16 {
		return 0, false
	}
	return int16(dt), true
}

func (dt DTInt64) MustInt16(args ...int16) int16 {
	var def int16
	if len(args) > 1 {
		log.Panicf("MustInt16() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt16 || dt < math.MinInt16 {
		return def
	}
	return int16(dt)
}

func (dt DTInt64) Int32() (int32, bool) {
	if dt > math.MaxInt32 || dt < math.MinInt32 {
		return 0, false
	}
	return int32(dt), true
}

func (dt DTInt64) MustInt32(args ...int32) int32 {
	var def int32
	if len(args) > 1 {
		log.Panicf("MustInt32() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if dt > math.MaxInt32 || dt < math.MinInt32 {
		return def
	}
	return int32(dt)
}

func (dt DTInt64) Int64() (int64, bool) {
	return int64(dt), true
}

func (dt DTInt64) MustInt64(args ...int64) int64 {
	if len(args) > 1 {
		log.Panicf("MustInt64() received too many arguments %d", len(args))
	}
	return int64(dt)
}

func (dt DTInt64) Float32() (float32, bool) {
	return float32(dt), true
}

func (dt DTInt64) MustFloat32(args ...float32) float32 {
	if len(args) > 1 {
		log.Panicf("MustFloat32() received too many arguments %d", len(args))
	}
	return float32(dt)
}

func (dt DTInt64) Float64() (float64, bool) {
	return float64(dt), true
}

func (dt DTInt64) MustFloat64(args ...float64) float64 {
	if len(args) > 1 {
		log.Panicf("MustFloat64() received too many arguments %d", len(args))
	}
	return float64(dt)
}
