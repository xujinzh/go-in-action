// 这个示例程序展示如何使用 io.Reader 和 io.Writer 接口写一个简单版本的 curl
/*
go build .
./writer1 https://github.com/filebrowser/filebrowser/releases/download/v2.63.17/linux-amd64-filebrowser.tar.gz linux-amd64-filebrowser.tar.gz
*/
package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/proxy"
)

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

// main 是应用程序的入口
func main() {
	// 1. 创建拨号器
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

	// 2. 将拨号器绑定到 Transport
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	// 3. 创建 Client
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	// 这里的 r 是一个响应，r.Body 是 io.Reader
	r, err := client.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// 创建文件来保存响应内容
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	// 使用 MultiWriter，这样就可以同时向文件和标准输出设备进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	// 读取响应的内容，并写到两个目的地
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
