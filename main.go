package main

import (
	"math"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type SineWave struct {
	freq       float64
	sampleRate beep.SampleRate
	phase      float64
}

// Stream generates the next samples of the sine wave
func (s *SineWave) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		sample := math.Sin(2 * math.Pi * s.phase)
		samples[i][0] = sample
		samples[i][1] = sample
		s.phase += s.freq / float64(s.sampleRate)
		if s.phase >= 1 {
			s.phase -= 1
		}
	}
	return len(samples), true
}

func (s *SineWave) Err() error {
	return nil
}

func main() {
	sampleRate := beep.SampleRate(10000)
	beepDuration := time.Second

	speaker.Init(sampleRate, sampleRate.N(time.Second/10))

	// Create a sine wave oscillator
	freq := 350.0 // Frequency of the beep
	sine := &SineWave{freq: freq, sampleRate: sampleRate}

	// Create a streamer that plays the sine wave for the specified duration
	streamer := beep.Take(sampleRate.N(beepDuration), sine)

	// Play the beep sound
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
