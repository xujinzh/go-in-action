// 这个示例程序展示如何在程序里造成竞争状态
// 实际上不希望出现这种情况
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // wg 用来等待程序结束
var counter int       // counter 是所有 goroutine 都要增加其值的变量

// main 是所有 Go 程序的入口
func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)
	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)
	// 等待 goroutine 结束
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

// incCounter 增加包里 counter 变量的值
func incCounter(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()
	for i := 0; i < 2; i++ {
		// 捕获 counter 的值
		value := counter
		// 当前 goroutine 从线程退出，并放回到队列
		/*
			调用 runtime 包的 Gosched 函数，用于将 goroutine 从当前线程退出，
			给其他 goroutine 运行的机会。在两次操作中间这样做的目的是强制调度器切换两个 goroutine，
			以便让竞争状态的效果变得更明显
		*/
		runtime.Gosched()
		// 增加本地 value 变量的值
		value++
		// 将该值保存回 counter
		counter = value
	}
}
