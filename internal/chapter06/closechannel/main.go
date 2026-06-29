/*
对于关闭的通道，如果缓冲区中还有值，那么还是可以读取这个值的。
*/
package main

import "fmt"

func main() {
	var c = make(chan int, 100)
	c <- 1
	c <- 2
	close(c)
	fmt.Println(<-c)
}
