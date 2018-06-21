package main

import (
	"fmt"
)

const limit = 10

func main() {
	// q := make(chan int)
	// c := gen(q)

	c, q := gen2()

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < limit; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func gen2() (chan int, chan int) {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < limit; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c, q
}
