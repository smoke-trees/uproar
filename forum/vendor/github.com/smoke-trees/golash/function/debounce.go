package function

import (
	"time"
)

// Debounced struct gives a debounced function which is called after certain intervals
type Debounced struct {
	function func()
	duration time.Duration
	stop     bool
	channel  chan bool
}

// Debounce returns a debounced function upon being passed a function f and a time duration t.
// On calling the returned (debounced function), the passed function is called repeatedly after fixed intervals
// Minimum time Interval should be 1 millisecond
func Debounce(f func(), t time.Duration) Debounced {

	if t < time.Millisecond*1 {
		panic("minimum value of t must be 1 millisecond")
	}

	debounced := Debounced{
		f, t, false, make(chan bool),
	}
	return debounced
}

// Flush instantly invokes debounced function
func (d *Debounced) Flush() {
	d.function()
}

// Call invokes the debounced Method so that the function is repeatedly called
func (d *Debounced) Call() {
	d.stop = false
	go func() {
		for !d.stop {
			select {
			case hg := <-d.channel:
				d.stop = hg
			default:
				d.stop = false
			}
			go d.function()
			time.Sleep(d.duration)
		}
	}()
}

// Cancel stops the debounced method
func (d *Debounced) Cancel() {
	d.channel <- true
}
