package series

import (
	"errors"
	"fmt"
	"github.com/hunknownz/blurr/internal/ndslice"
	"github.com/hunknownz/blurr/types"
)

type SeriesConfig struct {
	DType *types.DType
	Name  *string

	conFuncs []ConFunc
}

type ConFunc func(*SeriesConfig)

func WithName(name string) ConFunc {
	return func(sc *SeriesConfig) {
		tmp := name
		sc.Name = &tmp
	}
}

func WithDType(t types.DType) ConFunc {
	return func(sc *SeriesConfig) {
		tmp := t
		sc.DType = &tmp
	}
}

func (sc *SeriesConfig) Complete() {
	cfs := sc.conFuncs
	for i := range cfs {
		cfs[i](sc)
	}

	if sc.Name == nil {
		tmp := ""
		sc.Name = &tmp
	}
}

func (sc *SeriesConfig) New(values any) (s *Series, err error) {
	if sc.DType != nil && !sc.DType.IsValid() {
		err = errors.New(fmt.Sprintf("can't build series with dtype %s", *sc.DType))
	}

	if values == nil && sc.DType == nil {
		err = errors.New("can't build series without values and dtype")
		return
	}

	if values == nil && sc.DType != nil {
		nds, e := ndslice.MakeZeroSlice(*sc.DType, 0)
		if e != nil {
			err = fmt.Errorf("build series error: %w", e)
			return
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
		return
	}

	if values != nil && sc.DType == nil {
		switch v := values.(type) {
		case []int:
			nds, e := ndslice.NewSlice(v)
			if e != nil {
				err = fmt.Errorf("build series error: %w", e)
				return
			}
			s = &Series{
				name:    *sc.Name,
				ndslice: nds,
			}
		case []int8:
			nds, e := ndslice.NewSlice(v)
			if e != nil {
				err = fmt.Errorf("build series error: %w", e)
				return
			}
			s = &Series{
				name:    *sc.Name,
				ndslice: nds,
			}
		case []int16:
			nds, e := ndslice.NewSlice(v)
			if e != nil {
				err = fmt.Errorf("build series error: %w", e)
				return
			}
			s = &Series{
				name:    *sc.Name,
				ndslice: nds,
			}
		case []int32:
			nds, e := ndslice.NewSlice(v)
			if e != nil {
				err = fmt.Errorf("build series error: %w", e)
				return
			}
			s = &Series{
				name:    *sc.Name,
				ndslice: nds,
			}
		case []int64:
			nds, e := ndslice.NewSlice(v)
			if e != nil {
				err = fmt.Errorf("build series error: %w", e)
				return
			}
			s = &Series{
				name:    *sc.Name,
				ndslice: nds,
			}
		case []float32:
			nds, e := ndslice.NewSlice(v)
			if e != nil {
				err = fmt.Errorf("build series error: %w", e)
				return
			}
			s = &Series{
				name:    *sc.Name,
				ndslice: nds,
			}
		case []float64:
			nds, e := ndslice.NewSlice(v)
			if e != nil {
				err = fmt.Errorf("build series error: %w", e)
				return
			}
			s = &Series{
				name:    *sc.Name,
				ndslice: nds,
			}
		default:
			err = errors.New(fmt.Sprintf("can't build series with values with type %T", v))
			return
		}
		return
	}

	switch v := values.(type) {
	case []int:
		var nds *ndslice.RWNDSlice
		if *sc.DType == types.DInt {
			nds, err = ndslice.NewSlice(v)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
		} else {
			n := len(v)
			nds, err = ndslice.MakeZeroSlice(*sc.DType, n)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
			for i := 0; i < n; i++ {
				err = nds.ISet(i, types.DTInt(v[i]))
				if err != nil {
					err = fmt.Errorf("build series error: %w", err)
					return
				}
			}
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
	case []int8:
		var nds *ndslice.RWNDSlice
		if *sc.DType == types.DInt8 {
			nds, err = ndslice.NewSlice(v)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
		} else {
			n := len(v)
			nds, err = ndslice.MakeZeroSlice(*sc.DType, n)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
			for i := 0; i < n; i++ {
				err = nds.ISet(i, types.DTInt8(v[i]))
				if err != nil {
					err = fmt.Errorf("build series error: %w", err)
					return
				}
			}
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
	case []int16:
		var nds *ndslice.RWNDSlice
		if *sc.DType == types.DInt16 {
			nds, err = ndslice.NewSlice(v)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
		} else {
			n := len(v)
			nds, err = ndslice.MakeZeroSlice(*sc.DType, n)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
			for i := 0; i < n; i++ {
				err = nds.ISet(i, types.DTInt16(v[i]))
				if err != nil {
					err = fmt.Errorf("build series error: %w", err)
					return
				}
			}
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
	case []int32:
		var nds *ndslice.RWNDSlice
		if *sc.DType == types.DInt32 {
			nds, err = ndslice.NewSlice(v)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
		} else {
			n := len(v)
			nds, err = ndslice.MakeZeroSlice(*sc.DType, n)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
			for i := 0; i < n; i++ {
				err = nds.ISet(i, types.DTInt32(v[i]))
				if err != nil {
					err = fmt.Errorf("build series error: %w", err)
					return
				}
			}
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
	case []int64:
		var nds *ndslice.RWNDSlice
		if *sc.DType == types.DInt64 {
			nds, err = ndslice.NewSlice(v)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
		} else {
			n := len(v)
			nds, err = ndslice.MakeZeroSlice(*sc.DType, n)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
			for i := 0; i < n; i++ {
				err = nds.ISet(i, types.DTInt64(v[i]))
				if err != nil {
					err = fmt.Errorf("build series error: %w", err)
					return
				}
			}
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
	case []float32:
		var nds *ndslice.RWNDSlice
		if *sc.DType == types.DFloat32 {
			nds, err = ndslice.NewSlice(v)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
		} else {
			n := len(v)
			nds, err = ndslice.MakeZeroSlice(*sc.DType, n)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
			for i := 0; i < n; i++ {
				err = nds.ISet(i, types.DTFloat32(v[i]))
				if err != nil {
					err = fmt.Errorf("build series error: %w", err)
					return
				}
			}
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
	case []float64:
		var nds *ndslice.RWNDSlice
		if *sc.DType == types.DFloat64 {
			nds, err = ndslice.NewSlice(v)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
		} else {
			n := len(v)
			nds, err = ndslice.MakeZeroSlice(*sc.DType, n)
			if err != nil {
				err = fmt.Errorf("build series error: %w", err)
				return
			}
			for i := 0; i < n; i++ {
				err = nds.ISet(i, types.DTFloat64(v[i]))
				if err != nil {
					err = fmt.Errorf("build series error: %w", err)
					return
				}
			}
		}
		s = &Series{
			name:    *sc.Name,
			ndslice: nds,
		}
	default:
		err = errors.New(fmt.Sprintf("can't build series with values with type %T", v))
		return
	}
	return
}

func NewSeriesConfig(conFuncs ...ConFunc) (sc *SeriesConfig) {
	sc = new(SeriesConfig)
	sc.conFuncs = conFuncs
	return
}
