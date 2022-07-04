package internal

import (
	"errors"
	"fmt"
)

const (
	shift = 6
	mask  = 0x3f
)

type BitSet []uint64

func (bn *BitSet) GetBit(i int) bool {
	if i < 0 {
		return false
	}
	index := index(i)
	if index >= len(*bn) {
		return false
	}
	return (*bn)[index]&posVal(i) != 0
}

func (bn *BitSet) SetBit(i int, val bool) (oldVal bool, err error) {
	index := index(i)
	if index >= len(*bn) {
		err = errors.New(fmt.Sprintf("index out of slice with length %d", len(*bn)))
		return
	}

	oldVal = bn.GetBit(i)
	if !oldVal && val {
		(*bn)[index] |= posVal(i)
	} else if oldVal && !val {
		(*bn)[index] &^= posVal(i)
	}
	return
}

func index(i int) int {
	return i >> shift
}

func posVal(i int) uint64 {
	return 1 << uint(i&mask)
}

func NewBitSet(size int) *BitSet {
	bn := make(BitSet, size>>shift+1)
	return &bn
}
