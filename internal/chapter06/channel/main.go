package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// 接收者先启动，会被阻塞直到有发送
	go func() {
		fmt.Println("接收者：准备接收...")
		msg := <-ch
		fmt.Println("接收者：收到", msg)
	}()

	time.Sleep(1 * time.Second)
	str := "hello"
	fmt.Printf("发送者：发送 %s\n", str)
	ch <- str
	fmt.Println("发送者：发送完毕")
	time.Sleep(1 * time.Second)
}
