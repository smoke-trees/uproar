package function

// Afters struct for after functionality
type Afters struct {
	f func()
	n int
	i int
}

// After Returns a Afters struct with a Call and Reset method.
// Arguments to be passed :- function, n
// passed function gets invoked if the Call function on struct is called n or more times
// Reset function resets the count
func After(f func(), n int) Afters {
	if n > 0 {
		return Afters{
			f, n, 0,
		}
	}
	return Afters{
		f, 0, 0,
	}
}

// Call invokes the inner function
func (a *Afters) Call() {
	if a.i < a.n-1 {
		a.i++
	} else {
		a.f()
	}
}

// Reset resets the counter. Call function on struct has to be call n times more.
func (a *Afters) Reset() {
	a.i = 0
}
