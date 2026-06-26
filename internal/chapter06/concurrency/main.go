package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg用来等待协程完成
	var wg sync.WaitGroup
	// 等待2个协程
	wg.Add(2)

	fmt.Println("Start Goruntines")

	// 声明一个匿名函数，并启动一个协程
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		// 显示字母3次
		for i := 0; i < 3; i++ {
			for c := 'a'; c < 'a'+26; c++ {
				fmt.Printf("%c ", c)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for i := 0; i < 3; i++ {
			for c := 'A'; c < 'A'+26; c++ {
				fmt.Printf("%c ", c)
			}
		}
	}()

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
