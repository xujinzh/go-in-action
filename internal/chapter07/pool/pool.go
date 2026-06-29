// 包 pool 管理用户定义的一组资源
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全地在多个 goroutine 间共享的资源。
// 被管理的资源必须实现 io.Closer 接口
type Pool struct {
	// 互斥锁，保护 closed 状态标志
	m sync.Mutex
	// 缓冲通道，作为存储空闲资源的“池子”
	/*
		io.Closer 接口：这是 Go 的标准接口（包含一个 Close() error 方法）。
		这意味着该池子非常通用，任何实现了 Close() 的类型（如 *os.File, net.Conn 数据库连接）都能被这个池子管理。
	*/
	resources chan io.Closer
	// 函数指针，当池子空了，用来生成新资源的“工厂”
	factory func() (io.Closer, error)
	// 标志位，记录池子是否已经被关闭
	closed bool
}

// ErrPoolClosed 表示请求了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

// New 创建一个用来管理资源的池。
// 这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	/*
		懒加载设计：创建池子时，它没有立即生产任何资源，
		只是利用 make(chan io.Closer, size) 初始化了一个指定容量的空通道。

		工厂函数：传入的 fn 决定了以后如何创建具体的资源。
	*/
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	/*
		select + default 实现了非阻塞逻辑：如果 p.resources 里面有闲置资源，直接取走；
		如果里面没有，绝不阻塞等待，而是立刻调用 p.factory() 临时生产一个发给调用者。
	*/
	select {
	// 检查是否有空闲的资源
	case r, ok := <-p.resources: // 场景 A：池子里有现成的资源
		log.Println("Acquire:", "Shared Resource") // 池子关了，拒绝请求
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	// 因为没有空闲资源可用，所以提供了一个新资源
	default: // 场景 B：池子是空的（通道无数据）
		log.Println("Acquire:", "New Resource")
		return p.factory() // 动态创建一个新资源返回
	}
}

// Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和 Close 操作的安全
	p.m.Lock() // 加锁，防止在归还的同时，另外一个线程在执行 Close()
	defer p.m.Unlock()

	// 如果池子已经被关闭，销毁这个资源
	// 如果池子都关了，直接把归还的资源销毁
	if p.closed {
		r.Close()
		return
	}

	select {
	// 试图将这个资源放入队列
	case p.resources <- r: // 场景 A：池子还没满，放回通道里缓存起来
		log.Println("Release:", "In Queue")
	// 如果队列已满，则关闭这个资源
	default: // 场景 B：池子满了（排队满了）
		log.Println("Release:", "Closing")
		r.Close() // 放不下了，直接销毁该资源，防止内存泄露
	}
}

// Close 会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	// 保证本操作与 Release 操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	// 如果 pool 已经被关闭，什么也不做
	// 幂等性设计：防止重复关闭
	if p.closed {
		return
	}

	// 将池关闭
	p.closed = true

	// 在情况通道里的资源之前，将通道关闭
	// 如果不这样做，会发生死锁
	// 1. 先关闭通道（此时无法再写入，但缓冲区数据还在）
	close(p.resources)

	// 关闭资源
	// 2. 循环清空通道里的残余资源并逐个安全销毁
	for r := range p.resources {
		r.Close()
	}
}
