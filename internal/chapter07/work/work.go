/*
work 包的目的是展示如何使用无缓冲的通道来创建一个 goroutine 池，这些 goroutine 执行
并控制一组工作，让其并发执行。在这种情况下，使用无缓冲的通道要比随意指定一个缓冲区大
小的有缓冲的通道好，因为这个情况下既不需要一个工作队列，也不需要一组 goroutine 配合执
行。无缓冲的通道保证两个 goroutine 之间的数据交换。这种使用无缓冲的通道的方法允许使用
者知道什么时候 goroutine 池正在执行工作，而且如果池里的所有 goroutine 都忙，无法接受新的
工作的时候，也能及时通过通道来通知调用者。使用无缓冲的通道不会有工作在队列里丢失或者
卡住，所有工作都会被处理
*/
package work

import "sync"

// Worker必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 goroutine 停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
