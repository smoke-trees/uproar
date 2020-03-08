package function

import "time"

// Delay invokes the function f after t milliseconds
func Delay(f func(), t time.Duration) {

	if t < time.Millisecond*1 {
		panic("minimum value of t must be 1 millisecond")
	}

	go func() {
		time.Sleep(t)
		f()
	}()
}
