# GradBoostReg

[![GoDoc](https://godoc.org/github.com/siadat/gradboostreg/http?status.svg)](https://godoc.org/github.com/siadat/gradboostreg)
[![Build Status](https://travis-ci.org/siadat/gradboostreg.svg?branch=master)](https://travis-ci.org/siadat/gradboostreg)

Gradient Boosting Regressor in Go

## Install

    go get -u github.com/siadat/gradboostreg

## Example

### Problem
Let's say we have one predictor (also known as a feature) named X and we want to find the function
that predicts the value Y. Give the following 6 train samples:

    X:  0  1  2  3 4 5
    Y: 10 10 20 20 5 5

Predict Y for the following X values:

    X: 0.0 0.5 2.5 2.0 4.5

### Solution

```go
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
```
