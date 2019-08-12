package geiger // import "lukechampine.com/geiger"

import (
	"math"
	"math/rand"
	"runtime"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const sr = beep.SampleRate(44100)
const memSampleRate = 500 * time.Millisecond
const baseline = 160

func Count() {
	var p int
	speaker.Init(sr, sr.N(memSampleRate)/5)
	rng := rand.New(rand.NewSource(0x1337))
	sign := 1.0
	speaker.Play(beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		for i := range samples {
			if rng.Intn(int(sr)) < p {
				samples[i] = [2]float64{sign, sign}
				sign *= -1
			} else {
				samples[i] = [2]float64{0, 0}
			}
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
		// Allocs per second
		rate := float64(diff)/memSampleRate.Seconds() - baseline
		// If it becomes too high it becomes inaudible.
		// Limit at 3.6 Roentgen.
		rate = math.Min(rate, float64(sr)/3.6)
		speaker.Lock()
		p = int(rate)
		speaker.Unlock()
	}
}
