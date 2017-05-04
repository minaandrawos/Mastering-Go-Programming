package fanin

func fanin(chs ...<-chan int) <-chan int {
	out := make(chan int)
	for _, c := range chs {
		go registerChannel(c, out)
	}
	return out
}

func registerChannel(c <-chan int, out chan<- int) {
	for n := range c {
		out <- n
	}
}
