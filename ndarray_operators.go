package ndarray

import "math"

// ---------------------------------------------------------------
// Operator Functions all return a copy, original is not affected

func (this *Array) Cross(by *Array) *Array {
	n, m1 := this.Dims()
	m2, p := by.Dims()

	if m1 == m2 {
		dot := Zeros(n, p)
		// TODO
		return dot
	}

	return nil
}

func (this *Array) Mult(by *Array) *Array {
	t1, t2 := this.Type(), by.Type()

	if t1 == TYPE_VEC && t2 == TYPE_VEC {
		return this.PointwiseMult(by) // if both vectors, do a standard

	} else if t1 == TYPE_SCALER && t2 == TYPE_SCALER {
		return this.PointwiseMult(by) // if both

	} else if t1 == TYPE_MATRIX && t2 == TYPE_MATRIX {
		return Dot(this, by)

	} else if (t1 == TYPE_MATRIX || t1 == TYPE_VEC) && t2 == TYPE_SCALER {
		return this.Scale(by.AsScalar())

	} else if t1 == TYPE_SCALER && (t2 == TYPE_MATRIX || t2 == TYPE_VEC) {
		return by.Scale(this.AsScalar())

	}
	return nil
}

func (this *Array) PointwiseMult(by *Array) *Array {
	if len(this.Data) != len(by.Data) {
		if (this.IsVec() || this.IsMatrix()) && by.IsScalar() {
			return this.Scale(by.AsScalar())
		} else if (by.IsVec() || by.IsMatrix()) && this.IsScalar() {
			return by.Scale(this.AsScalar())
		}
		return nil
	}

	mult := Zeros(this.Dims())
	for i := range mult.Data {
		mult.Data[i] = this.Data[i] * by.Data[i]
	}
	return mult
}

func (this *Array) Add(by *Array) *Array {
	if len(this.Data) != len(by.Data) {
		return nil
	}

	ndarr := Zeros(this.Dims())
	for i := range ndarr.Data {
		ndarr.Data[i] = this.Data[i] + by.Data[i]
	}
	return ndarr
}

func (this *Array) Sub(by *Array) *Array {
	if len(this.Data) != len(by.Data) {
		return nil
	}

	ndarr := Zeros(this.Dims())
	for i := range ndarr.Data {
		ndarr.Data[i] = this.Data[i] - by.Data[i]
	}
	return ndarr
}

func (this *Array) Scale(k float64) *Array {
	scaled := Zeros(this.Dims())
	for i := range scaled.Data {
		scaled.Data[i] = this.Data[i] * k
	}
	return scaled
}

func (this *Array) Abs() float64 {
	return 0
}

func (this *Array) Distance() float64 {
	ls := this.PointwiseMult(this)
	return math.Sqrt(ls.Sum())
}

func (this *Array) Sum() float64 {
	sum := 0.0
	for _, v := range this.Data {
		sum += v
	}
	return sum
}
