package sample

import "fmt"

type DefaultSample struct {
	Xs map[string]float64
	Y  float64
}

func (s DefaultSample) GetY() float64 {
	return s.Y
}

func (s DefaultSample) Clone(newY float64) Sample {
	xs := make(map[string]float64)
	for k, v := range s.Xs {
		xs[k] = v
	}
	return DefaultSample{
		Xs: xs,
		Y:  newY,
	}
}

func (s DefaultSample) GetX(name string) float64 {
	if val, ok := s.Xs[name]; ok {
		return val
	}
	panic(fmt.Sprintf("unknown name %q", name))
}

func (s DefaultSample) GetNames() []string {
	keys := make([]string, len(s.Xs))
	i := 0
	for k := range s.Xs {
		keys[i] = k
		i++
	}
	return keys
}
