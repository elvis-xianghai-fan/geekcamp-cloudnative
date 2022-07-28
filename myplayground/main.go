package main

import (
	"fmt"
	"time"

	"github.com/elvis-xianghai-fan/geekcamp-cloudnative/playground/server"
)

func main() {
	s := server.New()

	divideByZero := &server.Work{
		Op:    func(a, b int) int { return a / b },
		A:     100,
		B:     0,
		Reply: make(chan int),
	}
	s <- divideByZero

	work1 := &server.Work{
		Op:    func(a, b int) int { return a + b },
		A:     100,
		B:     200,
		Reply: make(chan int),
	}
	s <- work1

	work2 := &server.Work{
		Op:    func(a, b int) int { return a - b },
		A:     100,
		B:     200,
		Reply: make(chan int),
	}
	s <- work2

	for {
		select {
		case res := <-divideByZero.Reply:
			fmt.Println(res)
		case res := <-work1.Reply:
			fmt.Println(res)
		case res := <-work2.Reply:
			fmt.Println(res)
		case <-time.After(time.Second):
			fmt.Println("No result in one second.")
		}
	}
}
