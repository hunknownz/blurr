package ndslice

import (
	"sync"
)

const (
	maxDims = 8
)

var (
	usedIntSlicePool [maxDims + 1]sync.Pool
	intSlicePoolOnce sync.Once
)

func initIntSlicePool() {
	for i := 0; i <= maxDims; i++ {
		size := i
		usedIntSlicePool[i].New = func() any {
			return make([]int, size)
		}
	}
}

func GetUsedIntSlice(size int) (used []int) {
	intSlicePoolOnce.Do(initIntSlicePool)
	if size > maxDims {
		return make([]int, size, size)
	}
	usedAny := usedIntSlicePool[size].Get()
	if usedAny == nil {
		return make([]int, size)
	}
	used = usedAny.([]int)
	return
}

func FreeUsedIntSlice(s []int) {
	if s == nil {
		return
	}
	size := cap(s)
	if size > maxDims {
		return
	}
	for i := 0; i < size; i++ {
		s[i] = 0
	}
	usedIntSlicePool[size].Put(s)
}
