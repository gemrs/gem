package id

// Generator returns a channel which is a producer of sequential identifiers
func Generator(startVal int) <-chan int {
	c := make(chan int)
	go func() {
		i := startVal
		for {
			c <- i
			i++
		}
	}()
	return c
}
