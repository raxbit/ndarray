package ndarray

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
)

const DegToRad = 0.017453292519943295769236907684886127134428718885417
const RadToDeg = 57.2957795131

// RealEpsilon gives an upper bound on the relative error due to rounding in floating point arithmetic.
var RealEpsilon float64

func init() {
	//RealEpsilon = 0.00001
	RealEpsilon = 0.5
}

// Vec is a 3 dimensional num
type Vec [3]float64

// String return a string representation of a num with 5 decimals per axis
func (v *Vec) String() string {
	return fmt.Sprintf("[%0.5f, %0.5f, %0.5f]", v[0], v[1], v[2])
}

var (
	UnitX = Vec{1, 0, 0}
	UnitY = Vec{0, 1, 0}
	UnitZ = Vec{0, 0, 1}
)

func V(x, y, z float64) *Vec {
	e := &Vec{}
	e[0] = x
	e[1] = y
	e[2] = z
	return e
}
func V2(x, y float64) *Vec {
	e := &Vec{}
	e[0] = x
	e[1] = y
	e[2] = 0
	return e
}

func ZV() *Vec {
	return &Vec{0, 0, 0}
}

func Z() *Vec {
	return &Vec{0, 0, 1}
}
func Y() *Vec {
	return &Vec{0, 1, 0}
}
func X() *Vec {
	return &Vec{1, 0, 0}
}

func (v *Vec) X() float64 {
	return v[0]
}
func (v *Vec) Y() float64 {
	return v[1]
}
func (v *Vec) Z() float64 {
	return v[2]
}

func (v *Vec) Clone() *Vec {
	return &Vec{
		v[0],
		v[1],
		v[2],
	}
}

func (v *Vec) Set(x, y, z float64) {
	v[0] = x
	v[1] = y
	v[2] = z
}

func (v *Vec) Copy(b *Vec) {
	v[0] = b[0]
	v[1] = b[1]
	v[2] = b[2]
}

func (v *Vec) Clear() *Vec {
	v[0] = 0
	v[1] = 0
	v[2] = 0
	return v
}

func (v *Vec) Add(b *Vec) *Vec {
	v[0] += b[0]
	v[1] += b[1]
	v[2] += b[2]
	return v
}

func (v *Vec) NewAdd(b *Vec) *Vec {
	return &Vec{
		v[0] + b[0],
		v[1] + b[1],
		v[2] + b[2],
	}
}

func (v *Vec) Sub(b *Vec) *Vec {
	v[0] -= b[0]
	v[1] -= b[1]
	v[2] -= b[2]
	return v
}

func (v *Vec) NewSub(b *Vec) *Vec {
	return &Vec{
		v[0] - b[0],
		v[1] - b[1],
		v[2] - b[2],
	}
}

func (v *Vec) AddScaledVector(b *Vec, t float64) *Vec {
	if math.IsNaN(t) {
		panic("scale value passed to Vec.AddScaledVector() is NaN")
	}
	v[0] += b[0] * t
	v[1] += b[1] * t
	v[2] += b[2] * t
	return v
}

func (v *Vec) Inverse() *Vec {
	v[0] = -v[0]
	v[1] = -v[1]
	v[2] = -v[2]
	return v
}

func (v *Vec) NewInverse() *Vec {
	return &Vec{
		-v[0],
		-v[1],
		-v[2],
	}
}

func (v *Vec) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v *Vec) SquareLength() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v *Vec) Distance(b *Vec) float64 {
	x := b.X() - v.X()
	y := b.Y() - v.Y()
	z := b.Z() - v.Z()
	return math.Sqrt(x*x + y*y + z*z)
}

func (v *Vec) Normalize() *Vec {
	length := v.Length()
	if length > 0 {
		v.Scale(1 / length)
	}
	return v
}

func (v *Vec) Scale(t float64) *Vec {
	v[0] *= t
	v[1] *= t
	v[2] *= t
	return v
}

func (v *Vec) NewScale(t float64) *Vec {
	return &Vec{
		v[0] * t,
		v[1] * t,
		v[2] * t,
	}
}

func (v *Vec) Dot(b *Vec) float64 {
	return v[0]*b[0] + v[1]*b[1] + v[2]*b[2]
}

// NewCross aka VectorProduct "%"
func (v *Vec) NewCross(b *Vec) *Vec {
	return &Vec{
		v[1]*b[2] - v[2]*b[1],
		v[2]*b[0] - v[0]*b[2],
		v[0]*b[1] - v[1]*b[0],
	}

}

// NewVectorProduct aka cross product
func (v *Vec) NewVectorProduct(b *Vec) *Vec {
	return v.NewCross(b)
}

// ScalarProduct calculates and returns the scalar product of this num
// with the given num.
func (v *Vec) ScalarProduct(b *Vec) float64 {
	return v[0]*b[0] + v[1]*b[1] + v[2]*b[2]
}

func (v *Vec) Mult(b *Vec) *Vec {
	v[0] *= b[0]
	v[1] *= b[1]
	v[2] *= b[2]
	return v
}

func (v *Vec) NewMult(b *Vec) *Vec {
	return &Vec{
		v[0] * b[0],
		v[1] * b[1],
		v[2] * b[2],
	}
}

func (v *Vec) Equals(b *Vec) bool {
	diff := math.Abs(v[0] - b[0])
	if diff > RealEpsilon {
		return false
	}
	diff = math.Abs(v[1] - b[1])
	if diff > RealEpsilon {
		return false
	}
	diff = math.Abs(v[2] - b[2])
	return diff < RealEpsilon
}

func (v *Vec) ToPixelVec() pixel.Vec {
	return pixel.V(v.X(), v.Y())
}

func (v *Vec) RotateAround(origin *Vec, rads float64) *Vec {
	//x_rot := (v.X())*math.Cos(rads) - (v.Y())*math.Sin(rads)
	//y_rot := (v.X())*math.Sin(rads) - (v.Y())*math.Cos(rads)
	//x_rot := (v.X()-origin.X())*math.Cos(rads) - (v.Y()-origin.Y())*math.Sin(rads) + origin.X()
	//y_rot := (v.X()-origin.X())*math.Cos(rads) + (v.Y()-origin.Y())*math.Sin(rads) + origin.Y()
	x_rot := math.Cos(rads)*(v.X()-origin.X()) - math.Sin(rads)*(v.Y()-origin.Y()) + origin.X()
	y_rot := math.Sin(rads)*(v.X()-origin.X()) - math.Cos(rads)*(v.Y()-origin.Y()) + origin.Y()

	v.Set(x_rot, y_rot, v.Z())
	return v
}

func (v *Vec) NewAddSub(b *Vec) *Vec {
	return &Vec{
		v[0] + b[0],
		v[1] - b[1],
		v[2] + b[2],
	}
}

func (v *Vec) NdArray() *Array {
	a := Zeros(len(*v))
	return a
}
