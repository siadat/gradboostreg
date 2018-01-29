package stat

import "math"

func StdDev(xs []float64) float64 {
	var ss, compensation float64
	mean := Mean(xs)
	for _, v := range xs {
		d := v - mean
		ss += d * d
		compensation += d
	}
	return math.Sqrt((ss - compensation*compensation/float64(len(xs))) / float64(len(xs)-1))
}

func Mean(s []float64) float64 {
	if len(s) == 0 {
		return 0.
	}
	return SumFloat64s(s) / float64(len(s))
}

func SumFloat64s(s []float64) float64 {
	sum := 0.0
	for i := range s {
		sum += s[i]
	}
	return sum
}

func AddScalar(s []float64, diff float64) []float64 {
	result := make([]float64, len(s))
	for i := range s {
		result[i] = s[i] + diff
	}
	return result
}
