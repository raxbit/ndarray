package ndarray

import "fmt"

const (
	TYPE_EMPTY  = 0
	TYPE_SCALER = 1
	TYPE_VEC    = 2
	TYPE_MATRIX = 4
	TYPE_N      = 8
)

func (this Array) String() string {
	switch this.Type() {
	case TYPE_EMPTY:
		return "[]"

	case TYPE_SCALER:
		return fmt.Sprintf("[1] %g", this.AsScalar())

	case TYPE_VEC:
		return fmt.Sprintf("[%d] (%v)", len(this.Data), this.Data)

	case TYPE_MATRIX:
		s := ""
		for i := 0; i < this.dims.Rows(); i++ {
			row := this.Row(i)
			s += fmt.Sprintf("%v,", row)
		}
		return fmt.Sprintf("[%d,%d] ([%s])", this.dims.Rows(), this.dims.Columns(), s[:len(s)-1])
	}
	return "NaN"

}
