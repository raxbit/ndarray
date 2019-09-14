package ndarray

import (
	"math"
)

func (d *Vec) IntersectsSphere(p0, p1 *Vec, r float64) float64 {

	//Sphere-> (x-sO.X)^2 + (y-sO.Y)^2 + (z-sO.Z)^2 = sR^2
	// Ray From->rO in Direction->v
	m := p0.NewSub(p1)
	b := m.Dot(d)
	c := m.Dot(m) - r*r
	if c > 0 && b > 0 {
		return -1 // Exit if r’s origin outside s (c > 0) and r pointing away from s (b > 0)
	}

	discr := b*b - c
	if discr < 0 {
		return -1 // A negative discriminant corresponds to ray missing sphere
	}

	t := -b - math.Sqrt(discr)
	if t < 0 {
		t = 0 // If t is negative, ray started inside sphere so clamp t to zero
	}
	//q := p0.NewAdd(&Vec{t * d, t * d, 0})

	return 1
}

func (v *Vec) IntersectsCircle() bool {
	return false
}

func (v *Vec) IntersectsSquare() bool {
	return false
}

func (v *Vec) IntersectsCylinder() bool {
	return false
}

func solveQuadratic(a, b, c float64) (ok bool, x0 float64, x1 float64) {
	discr := b*b - 4*a*c
	if discr < 0 {
		return false, 0, 0
	}

	if discr == 0 {
		x0 = -0.5 * b / a
		x1 = x0
	} else {
		q := 0.0
		if b > 0 {
			q = -0.5 * (b + math.Sqrt(discr))
		} else {
			q = -0.5 * (b - math.Sqrt(discr))
		}
		x0 = q / a
		x1 = c / q
	}

	if x0 > x1 {
		x1, x0 = x0, x1
	}

	return true, x0, x1
}
