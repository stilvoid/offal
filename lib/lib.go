package lib

import (
	"github.com/faiface/beep"
)

const RATE = 44100

var SampleRate beep.SampleRate

func init() {
	SampleRate = beep.SampleRate(RATE)
}
