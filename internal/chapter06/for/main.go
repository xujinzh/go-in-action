package main

import (
	"fmt"
	"runtime"
)

func main() {
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, "inches")

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
}
