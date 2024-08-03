package main

import (
	"math/rand/v2"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type Noise struct{}

func (noise Noise) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0] = rand.Float64()*2 - 1
		samples[i][1] = rand.Float64()*2 - 1
	}
	return len(samples), true
}

func (noise Noise) Err() error {
	return nil
}

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(Noise{})
	select {}
}
