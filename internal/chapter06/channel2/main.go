package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numGoroutines = 4 // 要使用的 goroutine 的数量
const taskLoad = 10     // 要处理的工作的数量

// wg 用来等待程序完成
var wg sync.WaitGroup
var r *rand.Rand

// init 初始化包，Go 语言运行时会在其他代码执行之前优先执行这个函数
func init() {
	source := rand.NewSource(3)
	r = rand.New(source)
}

// main 是所有 Go 程序的入口
func main() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作
	wg.Add(numGoroutines)
	for grWorker := 1; grWorker <= numGoroutines; grWorker++ {
		go worker(tasks, grWorker)
	}

	// 增加一组要完成的工作
	for t := 1; t <= taskLoad; t++ {
		tasks <- fmt.Sprintf("Task : %d", t)
	}

	// 当所有工作都处理完时关闭通道以便所有 goroutine 退出
	// 当通道关闭后，goroutine 依旧可以从通道接收数据，但是不能再向通道里发送数据。
	// 能够从已经关闭的通道接收数据这一点非常重要，因为这允许通
	// 道关闭后依旧能取出其中缓冲的全部值，而不会有数据丢失。从一个已经关闭且没有数据的通道
	// 里获取数据，总会立刻返回，并返回一个通道类型的零值
	close(tasks)
	// 等待所有工作完成
	wg.Wait()
}

// worker 作为 goroutine 启动来处理从有缓冲的通道传入的工作
func worker(tasks chan string, grWorker int) {
	// 通知函数已经返回
	defer wg.Done()

	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", grWorker)
			return
		}
		// 显示我们开始工作了
		fmt.Printf("Worker : %d : Started %s\n", grWorker, task)

		// 随机等一段时间来模拟工作
		sleep := r.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker : %d : Completed %s \n", grWorker, task)
	}
}
