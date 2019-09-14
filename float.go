package ndarray

func Max(base float64, compare float64) float64 {
	if (base >= compare) {
		return base
	}
	return  compare
}

func Min(base float64, compare float64) float64 {
	if (base <= compare) {
		return base
	}
	return  compare
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func Map(in, in_min, in_max, out_min, out_max float64) float64 {
	return out_min + ((out_max-out_min)/(in_max-in_min))*(in-in_min)
}