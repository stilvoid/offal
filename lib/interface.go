package lib

import (
	"github.com/faiface/beep"
)

type Operator func([]float64)

func (op Operator) Streamer() beep.StreamerFunc {
	return beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		output := make([]float64, len(samples))

		op(output)

		for i := range samples {
			samples[i][0] = (output[i] * 2) - 1
			samples[i][1] = (output[i] * 2) - 1
		}

		return len(samples), true
	})
}
