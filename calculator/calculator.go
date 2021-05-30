package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func anotherThread(c *Calculator) {
	for {
		// if our channel close then we recive false and we can undestand that we need to close output channel
		numInput, ok := <-(*c).Input
		if !ok {
			close((*c).Output)
			break
		} else {
			(*c).Output <- numInput * numInput
		}
	}
}
func (c *Calculator) Start() {
	// if we don't create a new Thread for calculator then our programm go to waiting data condition from channel calculator.input
	go anotherThread(c)
}
