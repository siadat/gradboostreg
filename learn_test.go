package gradboostreg_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/siadat/gradboostreg"
	"github.com/siadat/gradboostreg/sample"
)

func ExampleLearn() {
	trainSamples := []sample.Sample{
		sample.DefaultSample{Xs: map[string]float64{"X": 0}, Y: 10},
		sample.DefaultSample{Xs: map[string]float64{"X": 1}, Y: 10},
		sample.DefaultSample{Xs: map[string]float64{"X": 2}, Y: 20},
		sample.DefaultSample{Xs: map[string]float64{"X": 3}, Y: 20},
		sample.DefaultSample{Xs: map[string]float64{"X": 4}, Y: 5},
		sample.DefaultSample{Xs: map[string]float64{"X": 5}, Y: 5},
	}
	predictFunc := gradboostreg.Learn(trainSamples, 0.5, 10)

	testSamples := []sample.Sample{
		sample.DefaultSample{Xs: map[string]float64{"X": 0.0}},
		sample.DefaultSample{Xs: map[string]float64{"X": 0.5}},
		sample.DefaultSample{Xs: map[string]float64{"X": 2.5}},
		sample.DefaultSample{Xs: map[string]float64{"X": 2.0}},
		sample.DefaultSample{Xs: map[string]float64{"X": 4.5}},
	}

	for i := range testSamples {
		predictedY := predictFunc(testSamples[i])
		fmt.Printf("X=%.1f predictedY=%.1f\n", testSamples[i].GetX("X"), predictedY)
	}

	// Output:
	// X=0.0 predictedY=10.0
	// X=0.5 predictedY=10.0
	// X=2.5 predictedY=20.0
	// X=2.0 predictedY=20.0
	// X=4.5 predictedY=5.0
}

func TestLearn(t *testing.T) {
	samples := []sample.Sample{
		sample.DefaultSample{Xs: map[string]float64{"X": 0}, Y: 10},
		sample.DefaultSample{Xs: map[string]float64{"X": 1}, Y: 10},
		sample.DefaultSample{Xs: map[string]float64{"X": 2}, Y: 20},
		sample.DefaultSample{Xs: map[string]float64{"X": 3}, Y: 20},
		sample.DefaultSample{Xs: map[string]float64{"X": 4}, Y: 5},
		sample.DefaultSample{Xs: map[string]float64{"X": 5}, Y: 5},
	}
	predictFunc := gradboostreg.Learn(samples, 0.5, 30)
	checkSamples(t, samples, predictFunc, 0.00000001)

	predictFunc = gradboostreg.Learn(samples, 0.5, 500)
	checkSamples(t, samples, predictFunc, 0.00000001)
}

func TestLearnX2(t *testing.T) {
	var samples = []sample.Sample{}
	for i := -2.0; i < 2; i += 0.2 {
		x := i
		samples = append(samples, sample.DefaultSample{
			Xs: map[string]float64{"X": x},
			Y:  x * x,
		})
	}

	predictFunc := gradboostreg.Learn(samples, 0.5, 30)
	checkSamples(t, samples, predictFunc, 0.5)
}

func TestLearnX2Y2(t *testing.T) {
	var samples = []sample.Sample{}
	for i := -2.0; i < 2; i += 0.2 {
		for j := -2.0; j < 2; j += 0.2 {
			x, y := i, j
			samples = append(samples, sample.DefaultSample{
				Xs: map[string]float64{"X": x, "y": y},
				Y:  x*x + y*y,
			})
		}
	}

	predictFunc := gradboostreg.Learn(samples, 0.5, 50)
	checkSamples(t, samples, predictFunc, 1.0)
}

func TestLearnBanana(t *testing.T) {
	bananaFunc := func(x, y float64) float64 {
		a := 1.0
		b := 100.0
		return math.Pow(a-x, 2) + b*math.Pow(y-x*x, 2)
	}
	var samples = []sample.Sample{}
	for i := -2.0; i < 2; i += 0.2 {
		for j := -2.0; j < 2; j += 0.2 {
			x, y := i, j
			samples = append(samples, sample.DefaultSample{
				Xs: map[string]float64{"X": x, "y": y},
				Y:  bananaFunc(x, y),
			})
		}
	}

	predictFunc := gradboostreg.Learn(samples, 0.5, 20)
	checkSamples(t, samples, predictFunc, 500.0)
}

func checkSamples(t *testing.T, samples []sample.Sample, predictFunc gradboostreg.PredictFunc, allowedDiff float64) {
	for i := range samples {
		got, want := predictFunc(samples[i]), samples[i].GetY()
		if !(math.Abs(want-got) < allowedDiff) {
			t.Logf("i:%d want:%.2f got:%.2f diff:%v allowedDiff:%v",
				i, want, got, math.Abs(want-got), allowedDiff)
			t.FailNow()
		}
	}
}
