package main

import (
	"math/rand"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type Operator func([]float64, time.Time)

func (op Operator) Streamer() beep.StreamerFunc {
	return beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		output := make([]float64, len(samples))

		op(output, time.Now())

		for i := range samples {
			samples[i][0] = output[i]
			samples[i][1] = output[i]
		}

		return len(samples), true
	})
}

var Noise = Operator(func(samples []float64, t time.Time) {
	for i := range samples {
		samples[i] = rand.Float64()*2 - 1
	}
})

func Average(ops ...Operator) Operator {
	return Operator(func(samples []float64, t time.Time) {
		outputs := make([][]float64, len(ops))
		for i := range ops {
			// Init
			outputs[i] = make([]float64, len(samples))
			for j := range samples {
				outputs[i][j] = samples[j]
			}

			// Run
			ops[i](outputs[i], t)

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

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	//speaker.Play(Noise.Streamer())
	speaker.Play(Average(Noise, Noise).Streamer())
	select {}
}
