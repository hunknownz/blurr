package types

import (
	"fmt"
)

type GItem interface {
	GInteger | GFloat | ~string | ~bool
}

type GInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type GFloat interface {
	~float32 | ~float64
}

type DType uint8

const (
	DInvalid DType = iota
	DBool
	DInt
	DInt8
	DInt16
	DInt32
	DInt64
	DFloat32
	DFloat64
	DString
)

func (dt DType) String() string {
	if dt.IsValid() {
		return fmt.Sprintf("%s:%#v", mapDTypeString[dt], dt)
	}
	return fmt.Sprintf("%s:%#v", mapDTypeString[DInvalid], dt)
}

func (dt DType) IsValid() bool {
	return dt != DInvalid && dt.IsExists()
}

func (dt DType) IsExists() bool {
	return mapDTypeString[dt] != ""
}

/*
func (dt DType) Element(val any) (e Element, err error) {
	switch dt {
	case DInvalid:
		return &ElementInvalid{}
	case DInt:
		switch v := val.(type) {
		case int:
			e = ElementInt(v)
			return
		}
	}
	return
}
*/

const (
	InternalTypesNum = 10
	MaxTypesNum      = 256
)

var (
	mapDTypeString = [MaxTypesNum]string{
		"DInvalid", "DBool", "DInt",
		"DInt8", "DInt16", "DInt32",
		"DInt64", "DFloat32", "DFloat64",
		"DString",
	}
	mapNameDType = map[string]DType{
		"DInvalid": DInvalid,
		"DBool":    DBool,
		"DInt":     DInt,
		"DInt8":    DInt8,
		"DInt16":   DInt16,
		"DInt32":   DInt32,
		"DInt64":   DInt64,
		"DFloat32": DFloat32,
		"DFloat64": DFloat64,
		"DString":  DString,
	}
)

func GetDTypeByName(name string) (dtype DType, ok bool) {
	dtype, ok = mapNameDType[name]
	return
}
