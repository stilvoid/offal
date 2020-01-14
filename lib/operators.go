package lib

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
