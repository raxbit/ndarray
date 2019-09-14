package ndarray

import (
	"testing"
)

func TestZeros(t *testing.T) {
	a := Zeros(5)
	for _, v := range a.Data {
		if v != 0 {
			t.Error("Zeros", "0", v)
		}
	}
	rows, cols := a.Dims()
	if rows != 5 {
		t.Error("Zeros: rows ", "5", rows)
	}

	if cols != 1 {
		t.Error("Zeros: cols ", "1", cols)
	}
}

func TestOnes(t *testing.T) {
	a := Ones(3, 3)
	for _, v := range a.Data {
		if v != 1 {
			t.Error("Ones", "0", v)
		}
	}

	rows, cols := a.Dims()
	if rows != 3 {
		t.Error("Ones: rows ", "3", rows)
	}

	if cols != 3 {
		t.Error("Ones: cols ", "3", cols)
	}
}

func TestRandom(t *testing.T) {
	a := Random(2, 2)
	for _, v := range a.Data {
		if v == 0 {
			t.Error("Random", "0", v)
		}
	}

	rows, cols := a.Dims()
	if rows != 2 {
		t.Error("Random: rows ", "2", rows)
	}

	if cols != 2 {
		t.Error("Random: cols ", "2", cols)
	}
}

func TestNdArray_At(t *testing.T) {
	a := NewNdArray([]float64{1, 2, 3, 4}, 2, 2)

	if a.At(0, 0) != 1 {
		t.Error("At00", "1", a.At(0, 0))
	}
	if a.At(0, 1) != 2 {
		t.Error("At01", "2", a.At(0, 1))
	}
	if a.At(1, 0) != 3 {
		t.Error("At10", "3", a.At(1, 0))
	}
	if a.At(1, 1) != 4 {
		t.Error("At11", "4", a.At(1, 1))
	}
}

func TestNdArray_Set(t *testing.T) {
	a := Zeros(3, 3)
	a.Set(1, 2, 100)
	if a.At(1, 2) != 100 {
		t.Error("Set12", "100", a.At(1, 2))
	}

}

func TestNdArray_Dot(t *testing.T) {

}

func TestNdArray_Multiply(t *testing.T) {

}

func TestNdArray_T(t *testing.T) {
	a := NewNdArray([]float64{1, 2, 3, 4}, 2, 2)
	b := a.T()

	if a.Equal(b) {
		t.Error("Transponse", "false", a.Equal(b))
	}

	if b.At(0, 0) != 1 {
		t.Error("err00", b.At(0, 0))
	}
	if b.At(0, 1) != 3 {
		t.Error("err01", b.At(0, 1))
	}
	if b.At(1, 0) != 2 {
		t.Error("err10", b.At(1, 0))
	}
	if b.At(1, 1) != 4 {
		t.Error("err11", b.At(1, 1))
	}
}

func TestArrange(t *testing.T) {
	//a := Arrange(0.1, 1.0, 0.1)
	//log.Print("Arrange")
	//log.Print(a)
}
