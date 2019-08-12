package geiger // import "lukechampine.com/geiger"

import (
	"math"
	"runtime"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const sr = beep.SampleRate(44100)
const memSampleRate = 100 * time.Millisecond

func Count() {
	var freq, t float64
	step := sr.D(1).Seconds()
	speaker.Init(sr, sr.N(memSampleRate))
	speaker.Play(beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		for i := range samples {
			t += step
			a := math.Sin(2 * math.Pi * freq * t)
			samples[i] = [2]float64{a, a}
		}
		return len(samples), true
	}))

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	prevAllocs := ms.Mallocs
	for range time.Tick(memSampleRate) {
		runtime.ReadMemStats(&ms)
		diff := ms.Mallocs - prevAllocs
		prevAllocs = ms.Mallocs
		speaker.Lock()
		freq = float64(diff) / memSampleRate.Seconds()
		speaker.Unlock()
	}
}
