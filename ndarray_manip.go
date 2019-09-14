package ndarray

// Transpose
func (this *Array) T() *Array {
	r, c := this.Dims()
	trans := Zeros(c, r)

	for i := 0; i < r; i++ { // i
		for j := 0; j < c; j++ { // j
			trans.Set(j, i, this.At(i, j))
		}
	}
	return trans
}

// Apply a function over each element of the ndarray
func (this *Array) ForEach(fn func(x float64) float64) *Array {
	for i := range this.Data {
		this.Data[i] = fn(this.Data[i])
	}
	return this
}

// Like ForEach but creates and returns a copy
func (this Array) Activate(fn func(x float64) float64) *Array {
	ndarr := Zeros(this.Dims())
	for i := range ndarr.Data {
		ndarr.Data[i] = fn(this.Data[i])
	}
	return ndarr
}

func (this *Array) Normalize() *Array {
	length := this.Distance()
	if length > 0 {
		return this.Scale(1.0 / length)
	}
	return Copy(this)
}

func (this *Array) Diagonal() *Array {
	arr := Zeros(this.dims.Rows())
	for i := 0; i < this.dims.Rows(); i++ {
		for j := 0; j < this.dims.Columns(); j++ {
			if i == j {
				arr.Data[i] = this.At(i, j)
			}
		}
	}
	return arr
}

func (this *Array) Eye() *Array {
	arr := Zeros(this.Dims())
	for i := 0; i < this.dims.Rows(); i++ {
		for j := 0; j < this.dims.Columns(); j++ {
			if i == j {
				arr.Set(i, j, this.At(i, j))
			}
		}
	}
	return arr
}

func (this *Array) ArgMax(axis ...int) int {
	if len(axis) > 0 {
		return 0
	}

	max := 0.0
	index := 0
	for i, v := range this.Data {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}

func (this *Array) Flatten(order byte) *Array {
	if order == 0 {
		order = 'C'
	}

	switch order {
	case 'F':
		flat := make([]float64, 0)
		for i := 0; i < this.dims.Rows(); i++ {
			row := this.Row(i)
			flat = append(flat, row...)
		}
		return NewNdArray(flat, 1, len(this.Data))

	case 'C':
		flat := make([]float64, 0)
		for i := 0; i < this.dims.Columns(); i++ {
			col := this.Column(i)
			flat = append(flat, col...)
		}
		return NewNdArray(flat, 1, len(this.Data))

	}

	// Just return as is (should have same effect as 'F')
	return NewNdArray(this.Data, 1, len(this.Data))
}
