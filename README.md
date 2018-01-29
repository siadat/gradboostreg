    trainSamples := []sample.Sample{
    	sample.DefaultSample{Xs: map[string]float64{"x": 0}, Y: 10},
    	sample.DefaultSample{Xs: map[string]float64{"x": 1}, Y: 10},
    	sample.DefaultSample{Xs: map[string]float64{"x": 2}, Y: 20},
    	sample.DefaultSample{Xs: map[string]float64{"x": 3}, Y: 20},
    	sample.DefaultSample{Xs: map[string]float64{"x": 4}, Y: 5},
    	sample.DefaultSample{Xs: map[string]float64{"x": 5}, Y: 5},
    }
    predictFunc := gradboostreg.Learn(trainSamples, 0.5, 10)
    
    testSamples := []sample.Sample{
    	sample.DefaultSample{Xs: map[string]float64{"x": 0.0}, Y: 10},
    	sample.DefaultSample{Xs: map[string]float64{"x": 0.5}, Y: 10},
    	sample.DefaultSample{Xs: map[string]float64{"x": 2.5}, Y: 20},
    	sample.DefaultSample{Xs: map[string]float64{"x": 2.0}, Y: 20},
    	sample.DefaultSample{Xs: map[string]float64{"x": 4.5}, Y: 5},
    }
    
    for i := range testSamples {
    	predicted, actual := predictFunc(testSamples[i]), testSamples[i].GetY()
    	fmt.Printf("predicted=%.1f actual=%.1f\n", predicted, actual)
    }
    
    // Output:
    // predicted=10.0 actual=10.0
    // predicted=10.0 actual=10.0
    // predicted=20.0 actual=20.0
    // predicted=20.0 actual=20.0
    // predicted=5.0 actual=5.0
