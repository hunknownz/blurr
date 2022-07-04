package series

import (
	"fmt"
	"github.com/hunknownz/blurr/internal/ndslice"
	"github.com/hunknownz/blurr/types"
)

type Series struct {
	name    string
	ndslice *ndslice.RWNDSlice
}

func (s *Series) Size() int {
	return s.ndslice.Size()
}

func (s *Series) DType() types.DType {
	return s.ndslice.DType()
}

func (s *Series) Clone(deep bool) (ns *Series) {
	if !deep {
		ns = &Series{
			name:    s.name,
			ndslice: s.ndslice,
		}
		return
	}
	ns = &Series{
		name:    s.name,
		ndslice: s.ndslice.Clone(),
	}
	return
}

func New(vals interface{}, conFuncs ...ConFunc) (s *Series, err error) {
	sc := NewSeriesConfig(conFuncs...)
	sc.Complete()
	s, err = sc.New(vals)
	if err != nil {
		err = fmt.Errorf("new series error: %w", err)
		return
	}
	return
}
