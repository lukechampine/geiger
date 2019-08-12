package geiger

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
	"time"
)

// TestCount requires human ears.
func TestCount(t *testing.T) {
	go Count()
	tests := []struct {
		name string
		hz   int
		per  int
	}{
		{
			name: "baseline",
			hz:   1,
			per:  0,
		},
		{
			name: "5hz",
			hz:   5,
			per:  1,
		}, {
			name: "50hz",
			hz:   50,
			per:  1,
		}, {
			name: "500hz",
			hz:   50,
			per:  10,
		}, {
			name: "5000hz",
			hz:   50,
			per:  100,
		}, {
			name: "50000hz",
			hz:   50,
			per:  1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := time.After(5 * time.Second)
			every := time.NewTicker(time.Second / time.Duration(tt.hz))
			for {
				select {
				case <-done:
					return
				case <-every.C:
					for i := 0; i < tt.per; i++ {
						// alloc tmp
						tmp := bytes.Buffer{}
						io.Copy(ioutil.Discard, &tmp)
					}
				}
			}
		})
	}
}
