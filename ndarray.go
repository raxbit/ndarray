package ndarray

type Array struct {
	dims   *Shape
	Data   []float64
	Length int
}

// ---------------------------------------------------------------
func (this *Array) DataCopy() []float64 {
	d := make([]float64, len(this.Data))
	copy(d, this.Data)
	return d
}

func (this Array) Dims() (rows, columns int) {
	return this.dims.Rows(), this.dims.Columns()
}

func (this Array) Dim(index int) int {
	return this.dims.Dim(index)
}

func (this *Array) Row(i int) []float64 {
	row := make([]float64, this.dims.Columns())
	for j := range row {
		row[j] = this.At(i, j)
	}
	return row
}

func (this *Array) Column(j int) []float64 {
	columns := make([]float64, this.dims.Rows())
	for i := range columns {
		columns[j] = this.At(i, j)
	}
	return columns
}

func (this *Array) AsScalar() float64 {
	return this.Data[0]
}

func (this Array) At(r, c int) float64 {
	return this.Data[r*this.dims.Columns()+c]
}

func (this *Array) Set(r, c int, value float64) {
	this.Data[r*this.dims.Columns()+c] = value
}

// ---------------------------------------------------------------

func (this *Array) Equal(with *Array) bool {
	equal := true
	for i, v := range this.Data {
		if v != with.Data[i] {
			equal = false
			break
		}
	}
	return equal
}

func (this *Array) EqualDims(with *Array) bool {
	equal := true
	for i, v := range this.dims.dims {
		if v != with.Dim(i) {
			equal = false
			break
		}
	}
	return equal
}

func (this *Array) Inverse() *Array {
	return this.Scale(-1)
}

func (this Array) IsScalar() bool {
	r, c := this.Dims()
	if r == 1 && c == 1 {
		return true
	}
	return false
}

func (this Array) IsVec() bool {
	r, c := this.Dims()
	if r == 1 && c > 1 || r > 1 && c == 1 {
		return true
	}
	return false
}

func (this Array) IsMatrix() bool {
	r, c := this.Dims()
	if r > 1 && c > 1 {
		return true
	}
	return false
}

func (this Array) Type() int {
	r, c := this.Dims()
	if r == 0 || c == 0 {
		return TYPE_EMPTY
	} else if r == 1 && c == 1 {
		return TYPE_SCALER
	} else if (r == 1 && c > 1) || (r > 1 && c == 1) {
		return TYPE_VEC
	} else if r > 1 && c > 1 {
		return TYPE_MATRIX
	}
	return TYPE_N
}
