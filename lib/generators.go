package lib

import (
	"math"
	"math/rand"
)

var Noise = Operator(func(samples []float64) {
	for i := range samples {
		samples[i] = rand.Float64()*2 - 1
	}
})

func Sine(freq float64) Operator {
	t := 0.0

	return Operator(func(samples []float64) {
		for i := range samples {
			samples[i] = (math.Sin(math.Pi*2*freq*t) + 1) / 2
			t += SampleRate.D(1).Seconds()
		}
	})
}

func Square(freq float64) Operator {
	t := 0.0

	return Operator(func(samples []float64) {
		for i := range samples {
			y := math.Sin(math.Pi * 2 * freq * t)
			if y > 0 {
				samples[i] = 1
			} else {
				samples[i] = 0
			}
			t += SampleRate.D(1).Seconds()
		}
	})
}

func Saw(freq float64) Operator {
	t := 0.0

	return Operator(func(samples []float64) {
		for i := range samples {
			samples[i] = freq*t - math.Trunc(freq*t)
			t += SampleRate.D(1).Seconds()
		}
	})
}
