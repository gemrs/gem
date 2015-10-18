package id

// Generator returns a channel which is a producer of sequential identifiers
func Generator() <-chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			c <- i
			i++
		}
	}()
	return c
}
