package main

import (
	"time"

	"github.com/faiface/beep/speaker"
	"github.com/stilvoid/offal/lib"
)

func main() {
	speaker.Init(lib.SampleRate, lib.SampleRate.N(time.Second/10))

	//speaker.Play(Noise.Streamer())
	//speaker.Play(Average(Noise, Buzz).Streamer())
	//speaker.Play(Buzz.Streamer())
	//speaker.Play(Sine(440).Streamer())
	speaker.Play(lib.Saw(440).Streamer())
	/*
		speaker.Play(Multiply(
			Sine(449),
			Sine(331),
			Sine(211),
			Sine(103),
			Sine(13),
			Sine(7),
			Sine(3),
		).Streamer())
	*/
	select {}
}
