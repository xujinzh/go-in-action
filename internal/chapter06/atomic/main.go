// 这个示例程序展示如何使用 atomic 包来提供对数值类型的安全访问
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// counter 是所有 goroutine 都要增加其值的变量
var counter int64

// wg 用来等待程序结束
var wg sync.WaitGroup

// main 是所有 Go 程序的入口
func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)
	// 创建两个 goroutine
	go incCounter()
	go incCounter()
	// 等待 goroutine 结束
	wg.Wait()
	// 显示最终的值
	fmt.Println("Final counter:", counter)
}

// incCounter 增加包里 counter 变量的值
func incCounter() {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

	for range 2 {
		// 安全地对 counter 加 1
		/*
			这个函数会同步整型值的加法，方法是强制同一时刻只能有一个 goroutine 运行并完成这个加法操作
			当 goroutine 试图去调用任何原子函数时，这些 goroutine 都会自动根据所引用的变量做同步处理
		*/
		atomic.AddInt64(&counter, 1)
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}
}
