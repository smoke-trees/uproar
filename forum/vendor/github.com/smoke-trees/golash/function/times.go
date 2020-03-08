package function

// Counter uses the Counter.Call() function to invoke f once
type Counter struct {
	f     func(int)
	n     int
	count int
}

// Times Returns a Counter struct with a Call function that invokes Counter.f
func Times(f func(int), n int) Counter {
	c := Counter{
		f:     f,
		n:     n,
		count: 0,
	}
	return c
}

// Invoke Counter.f once
func (c *Counter) Call() {
	go func() {
		for c.count < c.n {
			c.f(c.count)
			c.count++
		}
	}()
}
