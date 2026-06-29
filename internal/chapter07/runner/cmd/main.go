// 这个示例程序演示如何使用通道来监视
// 程序运行的时间，以在程序运行时间过长时如何终止程序
package main

import (
	"log"
	"os"
	"time"

	"github.com/xujinzh/go-in-action/internal/chapter07/runner"
)

// timeout 规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

// main 是程序的入口
func main() {
	log.Println("Starting work.")

	// 为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	/*
		如果没有错误，任务就是按时执行完成的。
		如果执行超时，程序就会用错误码 1 终止。
		如果接收到中断信号，程序就会用错误码 2 终止。
		其他情况下，程序会使用错误码 0 正常终止。
	*/
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

// createTask 返回一个根据 id 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
