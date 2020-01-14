package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const RATE = 44100

var sampleRate beep.SampleRate

type Operator func([]float64)

func (op Operator) Streamer() beep.StreamerFunc {
	return beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		output := make([]float64, len(samples))

		op(output)

		for i := range samples {
			samples[i][0] = output[i]
			samples[i][1] = output[i]
		}

		return len(samples), true
	})
}

var Noise = Operator(func(samples []float64) {
	for i := range samples {
		samples[i] = rand.Float64()*2 - 1
	}
})

func Sine(freq float64) Operator {
	t := 0.0

	return Operator(func(samples []float64) {
		for i := range samples {
			y := math.Sin(math.Pi * 2 * freq * t)
			samples[i] = y
			t += sampleRate.D(1).Seconds()
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
			t += sampleRate.D(1).Seconds()
		}
	})
}

func Average(ops ...Operator) Operator {
	return Operator(func(samples []float64) {
		outputs := make([][]float64, len(ops))
		for i := range ops {
			// Init
			outputs[i] = make([]float64, len(samples))
			for j := range samples {
				outputs[i][j] = samples[j]
			}

			// Run
			ops[i](outputs[i])

			// Add
			for j := range samples {
				samples[j] += outputs[i][j]
			}
		}

		// Average
		for i := range samples {
			samples[i] /= float64(len(ops))
		}
	})
}

func Multiply(ops ...Operator) Operator {
	return Operator(func(samples []float64) {
		outputs := make([][]float64, len(ops))
		for i := range ops {
			// Init
			outputs[i] = make([]float64, len(samples))
			for j := range samples {
				outputs[i][j] = samples[j]
			}

			// Run
			ops[i](outputs[i])

			// Multiply
			for j := range samples {
				if i == 0 {
					samples[j] = outputs[i][j]
				} else {
					samples[j] *= outputs[i][j]
				}
			}
		}
	})
}

func main() {
	sampleRate = beep.SampleRate(RATE)
	speaker.Init(sampleRate, sampleRate.N(time.Second/2))
	//speaker.Play(Noise.Streamer())
	//speaker.Play(Average(Noise, Buzz).Streamer())
	//speaker.Play(Buzz.Streamer())
	//speaker.Play(Sine(440).Streamer())
	speaker.Play(Multiply(
		Sine(449),
		Sine(331),
		Sine(211),
		Sine(103),
		Sine(13),
		Sine(7),
		Sine(3),
	).Streamer())
	select {}
}
