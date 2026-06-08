package main

import (
	"log"
	"os"

	// 让 Go 语言对包做初始化操作，但是并不使用包里的标识符
	// 下划线让编译器接受这类导入，并且调用对应包内的所有代码文件里定义的 init 函数
	// 调用matchers 包中的 rss.go 代码文件里的 init 函数，注册 RSS 匹配器，以便后用
	_ "github.com/xujinzh/go-in-action/internal/chapter02/matchers"
	"github.com/xujinzh/go-in-action/internal/chapter02/search"
)

// init 在 main 之前调用
func init() {
	// 将日志输出到标准输出
	// 将标准库里日志类的输出，从默认的标准错误（stderr），设置为标准输出（stdout）设备
	log.SetOutput(os.Stdout)
}

// main 是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("president")
}
