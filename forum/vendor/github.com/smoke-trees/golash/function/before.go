package function

// Befores struct for before functionality
type Befores struct {
	f func()
	n int
	i int
}

// Before Returns a Befores struct with a Call and Reset method.
// Arguments to be passed :- function, n ( default value 1 for minimum 1 invocation)
// passed function gets invoked if the Call function in the struct is called n or less times
// Reset function resets the count
func Before(f func(), n int) Befores {
	if n > 0 {
		return Befores{
			f, n, 0,
		}
	}
	return Befores{
		f, 1, 0,
	}
}

// Call invokes the inner function
func (a *Befores) Call() {
	if a.i < a.n {
		a.f()
		a.i++
	} else {
		a.i++
	}
}

// Reset resets the counter. Call function on struct has to be call n times more.
func (a *Befores) Reset() {
	a.i = 0
}
