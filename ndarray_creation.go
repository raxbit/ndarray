package ndarray

import (
	"math/rand"

	"github.com/shopspring/decimal"
)

func NewNdArray(values []float64, rcd ...int) *Array {
	rows, cols := dimsToShape(rcd...)
	if len(values) != rows*cols {
		return nil // Invalid dimensions provided
	}

	a := Zeros(rcd...)
	a.Data = values
	a.Length = len(a.Data)
	return a
}

func NewRowVector(values []float64, length int) *Array {
	return NewNdArray(values, 1, length)
}

func HStack(a, b *Array) *Array {
	r, c := a.Dims()
	br, bc := b.Dims()

	//t := Copy(b)
	if r != br {
		return nil
	}

	m := Zeros(r, c+bc)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m.Set(i, j, a.At(i, j))
		}
		for j := 0; j < bc; j++ {
			m.Set(i, j+c, b.At(i, j))
		}
	}
	return m
}

func VStack(a, b *Array) *Array {
	r, c := a.Dims()
	br, bc := a.Dims()

	if c != bc {
		return nil
	}

	m := Zeros(r+br, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m.Set(i, j, a.At(i, j))
		}
	}
	for i := 0; i < br; i++ {
		for j := 0; j < c; j++ {
			m.Set(i+r, j, b.At(i, j))
		}
	}
	return m
}

func Empty() *Array {
	return &Array{dims: NewShape()}
}

func Scalar(value float64) *Array {
	s := Zeros(1)
	s.Data[0] = value
	return s
}

func VectorN(values ...float64) *Array {
	v := Zeros(len(values))
	v.Data = values
	return v
}

func Zeros(dims ...int) *Array {
	rows, cols := dimsToShape(dims...)
	a := &Array{
		Data:   make([]float64, rows*cols),
		dims:   NewShape(dims...),
		Length: rows*cols,
	}
	return a
}

func Arrange(start, stop, step float64) *Array {
	length := int((stop-start)/step) + 1
	arr := Zeros(length)

	curr := decimal.NewFromFloat(start)
	stride := decimal.NewFromFloat(step)
	for i := range arr.Data {
		arr.Data[i], _ = curr.Float64()
		curr = curr.Add(stride)
	}
	return arr
}

func ZerosLike(arr *Array) *Array {
	return Zeros(arr.Dims())
}

func Full(fill_value float64, dims ...int) *Array {
	a := Zeros(dims...)
	for i := range a.Data {
		a.Data[i] = fill_value
	}
	return a
}

func Ones(dims ...int) *Array {
	return Full(1, dims...)
}

func Random(dims ...int) *Array {
	a := Zeros(dims...)
	for i := range a.Data {
		a.Data[i] = rand.Float64()
	}
	return a
}

func Identity(dims ...int) *Array {
	a := Zeros(dims...)
	r, c := a.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == j {
				a.Set(i, j, 1.0)
			}
		}
	}
	return a
}

func Copy(arr *Array) *Array {
	cpy := ZerosLike(arr)
	for i, v := range arr.Data {
		cpy.Data[i] = v
	}
	cpy.Length = arr.Length
	return cpy
}

func Assign(lhs, rhs *Array) {
	copy(lhs.Data, rhs.Data)
}

func dimsToShape(dims ...int) (int, int) {
	rows, cols := 1, 1
	if len(dims) == 1 {
		rows = dims[0]
	} else if len(dims) == 2 {
		rows = dims[0]
		cols = dims[1]
	}
	return rows, cols
}
