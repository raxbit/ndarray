package ndarray

func VDot(this, by *Array) *Array {
	return Scalar(this.PointwiseMult(by).Sum())
}

func Dot(x, y *Array) *Array {
	t_this := x.Type()
	t_by := y.Type()

	if t_this == TYPE_VEC && t_by == TYPE_VEC {
		// if both vectors, do a standard
		return VDot(x, y)

	} else if t_this == TYPE_SCALER && t_by == TYPE_SCALER {
		return Scalar(x.AsScalar() * y.AsScalar())

	} else if t_this == TYPE_SCALER  {
		return nil // Not yet implemented

	} else if t_by == TYPE_SCALER {
		return nil // Not yet implemented
	}

	// We assume we have a MxM, VxM or MxV arrangement
	r1, c1 := x.Dims()
	r2, c2 := y.Dims()

	if c1 != r2 {
		return nil
	}

	out := Zeros(r1, c2)
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j ++ {
			out.Set(i,j, x.At(i,j) * y.At(j,i))
		}
	}

	return out
}



func (this *Array) Trace() float64 { // Sum of diagonals
	sum := 0.0
	for i := 0; i < this.dims.Rows(); i++ {
		for j := 0; j < this.dims.Columns(); j++ {
			if i == j {
				sum += this.At(i, j)
			}
		}
	}
	return sum
}

func (this *Array) Det() float64 {
	//TODO
	return 0
}

func vdot(v1, v2 []float64) float64 {
	total := 0.0
	for i, v := range v1 {
		total += v * v2[i]
	}
	return total
}
