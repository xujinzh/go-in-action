package main

import "fmt"

func main() {
	slice := make([]string, 5)
	fmt.Printf("slice: %v\n", slice)
	slice[0] = "3";
	fmt.Printf("slice: %v\n", slice)
}
