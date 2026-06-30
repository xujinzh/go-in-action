package main

import (
	"fmt"
	"log"
)

const (
	x1 = 1 << iota
	x2
	x3
	x4
	x5
	x6
)

func main() {
	fmt.Printf("x1: %v\n", x1)
	fmt.Printf("x2: %v\n", x2)
	fmt.Printf("x3: %v\n", x3)
	fmt.Printf("x4: %v\n", x4)
	fmt.Printf("x5: %v\n", x5)
	fmt.Printf("x6: %v\n", x6)

	fmt.Printf("log.Ldate: %v\n", log.Ldate)
	fmt.Printf("log.Llongfile: %v\n", log.Llongfile)

	fmt.Printf("log.Ldate type: %T\n", log.Ldate)
}
