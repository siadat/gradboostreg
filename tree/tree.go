package tree

import (
	"math"
	"runtime"
	"sort"
	"sync"

	"github.com/siadat/gradboostreg/sample"
	"github.com/siadat/gradboostreg/stat"
)

type Decider interface {
	Decide(sample.Sample) float64
}

type decisionTree struct {
	Threshold float64

	Feature string

	VarianceReduction float64

	Value1 float64
	Value2 float64

	Child1 *decisionTree
	Child2 *decisionTree
}

func (dt decisionTree) Decide(s sample.Sample) float64 {
	if s.GetX(dt.Feature) < dt.Threshold {
		if dt.Child1 == nil {
			return dt.Value1
		}
		return dt.Child1.Decide(s)
	}

	if dt.Child2 == nil {
		return dt.Value2
	}
	return dt.Child2.Decide(s)
}

type workStruct struct {
	x       float64
	leftw   float64
	rightw  float64
	feature string
	samples []sample.Sample
}

func newDecisionTree(stddev float64, feature string, x float64, leftw, rightw float64, samples []sample.Sample) decisionTree {
	ys1, ys2 := sample.SplitSampleYs(samples, feature, x)
	val1, val2 := stat.Mean(ys1), stat.Mean(ys2)
	var1diff, var2diff := 0.0, 0.0

	expectedVariance1 := stat.StdDev(stat.AddScalar(ys1, -val1))
	expectedVariance2 := stat.StdDev(stat.AddScalar(ys2, -val2))

	if math.IsNaN(expectedVariance1) {
		var1diff = 0.0
	} else {
		var1diff = stddev - expectedVariance1
	}
	if math.IsNaN(expectedVariance2) {
		var2diff = 0.0
	} else {
		var2diff = stddev - expectedVariance2
	}

	weightedVarianceReductions := []float64{leftw * var1diff, rightw * var2diff}

	return decisionTree{
		Value1:            val1,
		Value2:            val2,
		Threshold:         x,
		Feature:           feature,
		VarianceReduction: stat.SumFloat64s(weightedVarianceReductions),
	}
}

func newBestDecisionTree(samples []sample.Sample, alpha float64) decisionTree {
	if len(samples) == 0 {
		panic("len(samples)==0")
	}

	features := samples[0].GetNames()
	if len(features) == 0 {
		panic("no features")
	}

	var trees []decisionTree
	var treesLock sync.Mutex
	var stddev = stat.StdDev(sample.PluckSamplesYs(samples))
	var wg sync.WaitGroup
	var workChan = make(chan workStruct)

	for n := 0; n < runtime.NumCPU()*2; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for work := range workChan {
				tree := newDecisionTree(stddev, work.feature, work.x, work.leftw, work.rightw, work.samples)
				treesLock.Lock()
				trees = append(trees, tree)
				treesLock.Unlock()
			}
		}()
	}

	for _, feature := range features {
		xs := sample.PluckSamplesXs(samples, feature)
		sort.Float64s(xs)

		for i := range xs {
			workChan <- workStruct{
				x:       xs[i],
				leftw:   (float64(i) / float64(len(samples))),
				rightw:  (1.0 - float64(i)/float64(len(samples))),
				feature: feature,
				samples: samples,
			}
		}
	}

	close(workChan)
	wg.Wait()

	sort.Slice(trees, func(i, j int) bool {
		return trees[i].VarianceReduction < trees[j].VarianceReduction
	})
	return trees[len(trees)-1]
}

func NewDecisionTree(samples []sample.Sample, alpha float64) Decider {
	dtRoot := newBestDecisionTree(samples, alpha)
	samples1, samples2 := sample.SplitSamples(samples, dtRoot.Feature, dtRoot.Threshold)
	if len(samples1) > 0 {
		dtChild1 := newBestDecisionTree(samples1, alpha)
		dtRoot.Child1 = &dtChild1
	}
	if len(samples2) > 0 {
		dtChild2 := newBestDecisionTree(samples2, alpha)
		dtRoot.Child2 = &dtChild2
	}
	return dtRoot
}
