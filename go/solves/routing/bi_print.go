package main

import (
	"fmt"
)

func biprint() {
	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	doneA := make(chan struct{})
	doneB := make(chan struct{})
	chanA <- struct{}{}
	go func() {
		for i := 0; i < 10; i++ {
			<-chanA
			fmt.Println("1")
			chanB <- struct{}{}
		}
		doneA <- struct{}{}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-chanB
			fmt.Println("a")
			chanA <- struct{}{}
		}
		doneB <- struct{}{}
	}()
	<-doneA
	<-doneB
}

func main() {
	biprint()
}
