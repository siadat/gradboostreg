package gradboostreg

import (
	"github.com/siadat/gradboostreg/sample"
	"github.com/siadat/gradboostreg/stat"
	"github.com/siadat/gradboostreg/tree"
)

type PredictFunc func(sample.Sample) float64

func Learn(samples []sample.Sample, alpha float64, boostCount int) PredictFunc {
	if len(samples) == 0 {
		panic("Learn: len(samples)==0")
	}

	yMean := stat.Mean(sample.PluckSamplesYs(samples))
	dY := stat.AddScalar(sample.PluckSamplesYs(samples), -yMean)

	predictFuncs := make([]PredictFunc, boostCount)
	var samplesWithNewYs []sample.Sample
	for i := 0; i < boostCount; i++ {
		samplesWithNewYs = sample.SamplesWithNewYs(samples, dY)
		dt := tree.NewDecisionTree(samplesWithNewYs, alpha)
		predictFuncs[i] = func(s sample.Sample) float64 { return alpha * dt.Decide(s) }
		dY = newYs(dt, alpha, samplesWithNewYs)
	}

	return func(s sample.Sample) float64 {
		prediction := yMean
		for i := range predictFuncs {
			prediction = prediction + predictFuncs[i](s)
		}
		return prediction
	}
}

func newYs(dt tree.Decider, alpha float64, samples []sample.Sample) []float64 {
	diffs := make([]float64, len(samples))
	for i := range samples {
		diffs[i] = samples[i].GetY() - alpha*dt.Decide(samples[i])
	}
	return diffs
}
