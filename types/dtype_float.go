//go:generate go run gen_dtype.go -type DTFloat*
package types

import (
	"log"
	"math"
)

type DTFloat32 float32

func (dt DTFloat32) Int() (int, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxIntFloat := float64(math.MaxInt)
	if floatGreater(float64(dt), maxIntFloat) || floatLess(float64(dt), -maxIntFloat) {
		return 0, false
	}
	return int(dt), true
}

func (dt DTFloat32) MustInt(args ...int) int {
	var def int
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxIntFloat := float64(math.MaxInt)
	if floatGreater(float64(dt), maxIntFloat) || floatLess(float64(dt), -maxIntFloat) {
		return def
	}
	return int(dt)
}

func (dt DTFloat32) Int8() (int8, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt8Float := float64(math.MaxInt8)
	if floatGreater(float64(dt), maxInt8Float) || floatLess(float64(dt), -maxInt8Float) {
		return 0, false
	}
	return int8(dt), true
}

func (dt DTFloat32) MustInt8(args ...int8) int8 {
	var def int8
	if len(args) > 1 {
		log.Panicf("MustInt8() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxInt8Float := float64(math.MaxInt8)
	if floatGreater(float64(dt), maxInt8Float) || floatLess(float64(dt), -maxInt8Float) {
		return def
	}
	return int8(dt)
}

func (dt DTFloat32) Int16() (int16, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt16Float := float64(math.MaxInt16)
	if floatGreater(float64(dt), maxInt16Float) || floatLess(float64(dt), -maxInt16Float) {
		return 0, false
	}
	return int16(dt), true
}

func (dt DTFloat32) MustInt16(args ...int16) int16 {
	var def int16
	if len(args) > 1 {
		log.Panicf("MustInt16() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxInt16Float := float64(math.MaxInt16)
	if floatGreater(float64(dt), maxInt16Float) || floatLess(float64(dt), -maxInt16Float) {
		return def
	}
	return int16(dt)
}

func (dt DTFloat32) Int32() (int32, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt32Float := float64(math.MaxInt32)
	if floatGreater(float64(dt), maxInt32Float) || floatLess(float64(dt), -maxInt32Float) {
		return 0, false
	}
	return int32(dt), true
}

func (dt DTFloat32) MustInt32(args ...int32) int32 {
	var def int32
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxInt32Float := float64(math.MaxInt32)
	if floatGreater(float64(dt), maxInt32Float) || floatLess(float64(dt), -maxInt32Float) {
		return def
	}
	return int32(dt)
}

func (dt DTFloat32) Int64() (int64, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt64Float := float64(math.MaxInt64)
	if floatGreater(float64(dt), maxInt64Float) || floatLess(float64(dt), -maxInt64Float) {
		return 0, false
	}
	return int64(dt), true
}

func (dt DTFloat32) MustInt64(args ...int64) int64 {
	var def int64
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxIntFloat := float64(math.MaxInt)
	if floatGreater(float64(dt), maxIntFloat) || floatLess(float64(dt), -maxIntFloat) {
		return def
	}
	return int64(dt)
}

func (dt DTFloat32) Float32() (float32, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	return float32(dt), true
}

func (dt DTFloat32) MustFloat32(args ...float32) float32 {
	var def float32
	if len(args) > 1 {
		log.Panicf("MustFloat32() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	return float32(dt)
}

func (dt DTFloat32) Float64() (float64, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	return float64(dt), true
}

func (dt DTFloat32) MustFloat64(args ...float64) float64 {
	var def float64
	if len(args) > 1 {
		log.Panicf("MustFloat32() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	return float64(dt)
}

type DTFloat64 float64

func (dt DTFloat64) Int() (int, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxIntFloat := float64(math.MaxInt)
	if floatGreater(float64(dt), maxIntFloat) || floatLess(float64(dt), -maxIntFloat) {
		return 0, false
	}
	return int(dt), true
}

func (dt DTFloat64) MustInt(args ...int) int {
	var def int
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxIntFloat := float64(math.MaxInt)
	if floatGreater(float64(dt), maxIntFloat) || floatLess(float64(dt), -maxIntFloat) {
		return def
	}
	return int(dt)
}

func (dt DTFloat64) Int8() (int8, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt8Float := float64(math.MaxInt8)
	if floatGreater(float64(dt), maxInt8Float) || floatLess(float64(dt), -maxInt8Float) {
		return 0, false
	}
	return int8(dt), true
}

func (dt DTFloat64) MustInt8(args ...int8) int8 {
	var def int8
	if len(args) > 1 {
		log.Panicf("MustInt8() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxInt8Float := float64(math.MaxInt8)
	if floatGreater(float64(dt), maxInt8Float) || floatLess(float64(dt), -maxInt8Float) {
		return def
	}
	return int8(dt)
}

func (dt DTFloat64) Int16() (int16, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt16Float := float64(math.MaxInt16)
	if floatGreater(float64(dt), maxInt16Float) || floatLess(float64(dt), -maxInt16Float) {
		return 0, false
	}
	return int16(dt), true
}

func (dt DTFloat64) MustInt16(args ...int16) int16 {
	var def int16
	if len(args) > 1 {
		log.Panicf("MustInt16() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxInt16Float := float64(math.MaxInt16)
	if floatGreater(float64(dt), maxInt16Float) || floatLess(float64(dt), -maxInt16Float) {
		return def
	}
	return int16(dt)
}

func (dt DTFloat64) Int32() (int32, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt32Float := float64(math.MaxInt32)
	if floatGreater(float64(dt), maxInt32Float) || floatLess(float64(dt), -maxInt32Float) {
		return 0, false
	}
	return int32(dt), true
}

func (dt DTFloat64) MustInt32(args ...int32) int32 {
	var def int32
	if len(args) > 1 {
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxInt32Float := float64(math.MaxInt32)
	if floatGreater(float64(dt), maxInt32Float) || floatLess(float64(dt), -maxInt32Float) {
		return def
	}
	return int32(dt)
}

func (dt DTFloat64) Int64() (int64, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	maxInt64Float := float64(math.MaxInt64)
	if floatGreater(float64(dt), maxInt64Float) || floatLess(float64(dt), -maxInt64Float) {
		return 0, false
	}
	return int64(dt), true
}

func (dt DTFloat64) MustInt64(args ...int64) int64 {
	var def int64
	if len(args) > 1 {
		log.Panicf("MustInt64() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	maxInt64Float := float64(math.MaxInt64)
	if floatGreater(float64(dt), maxInt64Float) || floatLess(float64(dt), -maxInt64Float) {
		return def
	}
	return int64(dt)
}

func (dt DTFloat64) Float32() (float32, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	if floatGreater(float64(dt), math.MaxFloat32) || floatLess(float64(dt), -math.MaxFloat32) {
		return 0, false
	}
	return float32(dt), true
}

func (dt DTFloat64) MustFloat32(args ...float32) float32 {
	var def float32
	if len(args) > 1 {
		log.Panicf("MustInt64() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	if floatGreater(float64(dt), math.MaxFloat32) || floatLess(float64(dt), -math.MaxFloat32) {
		return def
	}
	return float32(dt)
}

func (dt DTFloat64) Float64() (float64, bool) {
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return 0, false
	}
	if math.IsNaN(float64(dt)) {
		return 0, false
	}
	return float64(dt), true
}

func (dt DTFloat64) MustFloat64(args ...float64) float64 {
	var def float64
	if len(args) > 1 {
		log.Panicf("MustFloat32() received too many arguments %d", len(args))
	}
	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	}
	if math.IsInf(float64(dt), 1) || math.IsInf(float64(dt), -1) {
		return def
	}
	if math.IsNaN(float64(dt)) {
		return def
	}
	return float64(dt)
}
