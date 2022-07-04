package ndslice

import (
	"errors"
	"fmt"
	"github.com/hunknownz/blurr/internal"
	"github.com/hunknownz/blurr/types"
	"reflect"
	"sync"
)

type RWNDSlice struct {
	RNDSlice
}

func (nds *RWNDSlice) ISet(i int, elem types.Element) (err error) {
	nds.mu.Lock()
	defer nds.mu.Unlock()

	var isNaN bool
	isNaN, err = nds.array.ISet(i, elem)
	if err != nil {
		err = fmt.Errorf("ndslice set error: %w", err)
		return
	}
	if isNaN {
		_, err = nds.iSetNaN(i, isNaN)
		if err != nil {
			err = fmt.Errorf("ndslice set nan error: %w", err)
			return
		}
	}
	return
}

func (nds *RWNDSlice) RWReshape(dims ...int) (nnds *RWNDSlice, err error) {
	if len(dims) == 0 {
		err = errors.New("total size of new array must be unchanged")
		return
	}

	if getDimsSize(dims...) != nds.array.Size() {
		err = errors.New("total size of new array must be unchanged")
		return
	}
	nnds = new(RWNDSlice)
	nnds.setShape(dims...)
	nnds.array = nds.array
	return
}

type RNDSlice struct {
	DMeta

	mu       sync.RWMutex
	array    types.Elements
	nanOnce  sync.Once
	nanArray *internal.BitSet
}

func (nds *RNDSlice) iSetNaN(i int, v bool) (oldNaNVal bool, err error) {
	nds.nanOnce.Do(func() {
		nds.nanArray = internal.NewBitSet(nds.Size())
	})
	oldNaNVal, err = nds.nanArray.SetBit(i, v)
	return
}

func (nds *RNDSlice) IGet(i int) (e types.Element, err error) {
	if i > nds.Size() {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", nds.Size()))
	}
	e, err = nds.array.IGet(i)
	if err != nil {
		err = fmt.Errorf("ndslice iget error: %w", err)
		return
	}
	return
}

func (nds *RNDSlice) Size() int {
	return nds.array.Size()
}

func (nds *RNDSlice) DType() types.DType {
	return nds.array.DType()
}

func (nds *RNDSlice) Clone() (nnds *RWNDSlice) {
	nds.mu.RLock()
	defer nds.mu.RUnlock()

	nnds = new(RWNDSlice)
	nnds.setShape(nds.shape...)
	nnds.array = nds.array.Clone()
	return
}

func (nds *RNDSlice) RReshape(dims ...int) (nnds *RNDSlice, err error) {
	if len(dims) == 0 {
		err = errors.New("total size of new array must be unchanged")
		return
	}

	if getDimsSize(dims...) != nds.array.Size() {
		err = errors.New("total size of new array must be unchanged")
		return
	}
	nnds = new(RNDSlice)
	nnds.setShape(dims...)
	nnds.array = nds.array
	return
}

func MakeZeroSlice(dtype types.DType, n int) (nds *RWNDSlice, err error) {
	if !dtype.IsValid() {
		err = errors.New(fmt.Sprintf("can't build ndslice with dtype %s", dtype))
	}

	dims := []int{n}
	nds = new(RWNDSlice)
	nds.setShape(dims...)
	switch dtype {
	case types.DInt:
		newArray := make(types.ElementsInt, n)
		nds.array = &newArray
	case types.DInt8:
		newArray := make(types.ElementsInt8, n)
		nds.array = &newArray
	case types.DInt16:
		newArray := make(types.ElementsInt16, n)
		nds.array = &newArray
	case types.DInt32:
		newArray := make(types.ElementsInt32, n)
		nds.array = &newArray
	case types.DInt64:
		newArray := make(types.ElementsInt64, n)
		nds.array = &newArray
	case types.DFloat32:
		newArray := make(types.ElementsFloat32, n)
		nds.array = &newArray
	case types.DFloat64:
		newArray := make(types.ElementsFloat64, n)
		nds.array = &newArray
	default:
		return
	}
	return
}

func NewSlice(vals any, dims ...int) (nds *RWNDSlice, err error) {
	if vals == nil {
		err = errors.New("input values must not be nil")
		return
	}

	t := reflect.TypeOf(vals)
	k := t.Kind()
	if k != reflect.Slice && k != reflect.Array {
		err = errors.New("input values should be slice or array")
		return
	}
	arrayLen := reflect.ValueOf(vals).Len()

	if len(dims) == 0 {
		dims = []int{arrayLen}
	}

	if getDimsSize(dims...) != arrayLen {
		err = errors.New("total size of new array must be unchanged")
		return
	}

	nds = new(RWNDSlice)
	nds.setShape(dims...)
	switch val := vals.(type) {
	case []int:
		newArray := make(types.ElementsInt, arrayLen)
		for i := 0; i < arrayLen; i++ {
			newArray[i] = types.DTInt(val[i])
		}
		nds.array = &newArray
	case []int8:
		newArray := make(types.ElementsInt8, arrayLen)
		for i := 0; i < arrayLen; i++ {
			newArray[i] = types.DTInt8(val[i])
		}
		nds.array = &newArray
	case []int16:
		newArray := make(types.ElementsInt16, arrayLen)
		for i := 0; i < arrayLen; i++ {
			newArray[i] = types.DTInt16(val[i])
		}
		nds.array = &newArray
	case []int32:
		newArray := make(types.ElementsInt32, arrayLen)
		for i := 0; i < arrayLen; i++ {
			newArray[i] = types.DTInt32(val[i])
		}
		nds.array = &newArray
	case []int64:
		newArray := make(types.ElementsInt64, arrayLen)
		for i := 0; i < arrayLen; i++ {
			newArray[i] = types.DTInt64(val[i])
		}
		nds.array = &newArray
	case []float32:
		newArray := make(types.ElementsFloat32, arrayLen)
		for i := 0; i < arrayLen; i++ {
			newArray[i] = types.DTFloat32(val[i])
		}
		nds.array = &newArray
	case []float64:
		newArray := make(types.ElementsFloat64, arrayLen)
		for i := 0; i < arrayLen; i++ {
			newArray[i] = types.DTFloat64(val[i])
		}
		nds.array = &newArray
	default:
		err = errors.New(fmt.Sprintf("can't build ndslice with type %T", val))
		return
	}

	return
}

func getDimsSize(dims ...int) (total int) {
	total = 1
	for i := 0; i < len(dims); i++ {
		total *= dims[i]
	}
	return
}
