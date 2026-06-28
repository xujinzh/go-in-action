package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var r *rand.Rand

func init() {
	source := rand.NewSource(42)
	r = rand.New(source)
}

func main() {

	court := make(chan int)
	wg.Add(2)

	go player("liming", court)
	go player("xiaohua", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := r.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}
