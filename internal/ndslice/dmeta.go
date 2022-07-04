package ndslice

type DMeta struct {
	shape   dShape
	strides []int
}

func (m *DMeta) setShape(dims ...int) {
	used := GetUsedIntSlice(len(dims))
	copy(used, dims)
	m.shape = dShape(used)

	m.strides = m.getStrides(dims)
}

func (m *DMeta) getStrides(dims []int) (strides []int) {
	used := GetUsedIntSlice(len(dims))
	cnt := 1
	for i := len(dims) - 1; i >= 0; i-- {
		used[i] = cnt
		cnt *= dims[i]
	}
	strides = used
	return
}

func (m *DMeta) NDim() int {
	return len(m.shape)
}
