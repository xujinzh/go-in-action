package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

func main() {
	// 1. 创建 SOCKS5 拨号器
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		panic(err)
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

	// 4. 发起请求
	resp, err := client.Get("https://www.cip.cc/")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("当前SOCKS请求使用的IP:", string(body))
}
