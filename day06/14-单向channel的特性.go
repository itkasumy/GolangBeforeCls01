package main

func main() {
	ch := make(chan int)

	var writeCh chan <- int = ch
	writeCh <- 666

	var readCh <-chan int = ch
	<-readCh
}
