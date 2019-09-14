package ndarray

type Shape struct {
	dims []int
}

func NewShape(dims ...int) *Shape {
	s := Shape{
		dims: make([]int, 3),
	}
	if len(dims) == 1 {
		s.dims[0] = dims[0]
		s.dims[1] = 1
	} else if len(dims) == 2 {
		s.dims[0] = dims[0]
		s.dims[1] = dims[1]
	} else if len(dims) == 3 {
		copy(s.dims, dims)
	}
	return &s
}

func (s Shape) Rows() int {
	return s.dims[0]
}

func (s Shape) Columns() int {
	return s.dims[1]
}

func (s Shape) Depth() int {
	return s.dims[2]
}

func (s Shape) Dims() []int {
	return s.dims
}

func (s Shape) Dim(index int) int {
	return s.dims[index]
}

func (s *Shape) Resize(dims ...int) {
	s = NewShape(dims...)
}
