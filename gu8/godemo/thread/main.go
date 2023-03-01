package main

import "fmt"

func printA(chanA, chanB, doneA chan struct{}) {
	for i := 0; i < 50; i++ {
		<-chanA
		fmt.Println("A")
		chanB <- struct{}{}
	}
	close(doneA)
}

func printB(chanA, chanB, doneB chan struct{}) {
	for i := 0; i < 50; i++ {
		<-chanB
		fmt.Println("B")
		chanA <- struct{}{}
	}
	close(doneB)
}

func main() {
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	doneChanA := make(chan struct{})
	doneChanB := make(chan struct{})
	chanA <- struct{}{}
	go printA(chanA, chanB, doneChanA)
	go printB(chanA, chanB, doneChanB)
	<-doneChanA
	<-doneChanB
}
