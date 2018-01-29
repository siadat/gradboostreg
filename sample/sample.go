package sample

type Sample interface {
	GetY() float64
	GetX(string) float64
	GetNames() []string
	Clone(newY float64) Sample
}

func SplitSampleYs(samples []Sample, feature string, threshold float64) ([]float64, []float64) {
	ys1 := make([]float64, 0, len(samples))
	ys2 := make([]float64, 0, len(samples))
	for _, s := range samples {
		if s.GetX(feature) < threshold {
			ys1 = append(ys1, s.GetY())
		} else {
			ys2 = append(ys2, s.GetY())
		}
	}
	return ys1, ys2
}

func SplitSamples(samples []Sample, feature string, threshold float64) (s1 []Sample, s2 []Sample) {
	for _, s := range samples {
		if s.GetX(feature) < threshold {
			s1 = append(s1, s)
		} else {
			s2 = append(s2, s)
		}
	}
	return s1, s2
}

func SamplesWithNewYs(samples []Sample, ys []float64) []Sample {
	if len(samples) != len(ys) {
		panic("len(samples) != len(ys)")
	}
	results := make([]Sample, len(samples))
	for i := range samples {
		results[i] = samples[i].Clone(ys[i])
	}
	return results
}

func PluckSamplesXs(samples []Sample, name string) []float64 {
	values := make([]float64, len(samples))
	for i := range samples {
		values[i] = samples[i].GetX(name)
	}
	return values
}

func PluckSamplesYs(samples []Sample) []float64 {
	values := make([]float64, len(samples))
	for i := range samples {
		values[i] = samples[i].GetY()
	}
	return values
}
