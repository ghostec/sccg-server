package math

import "time"

// WithFrameInterval func
func WithFrameInterval(fps float64, f func() error) error {
	start := time.Now()
	err := f()
	time.Sleep(time.Duration(1.0/fps)*time.Millisecond - (time.Now().Sub(start)))
	return err
}
