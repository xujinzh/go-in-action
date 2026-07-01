package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// 1. 解析代理地址
	proxyURL, err := url.Parse("http://127.0.0.1:8118")
	if err != nil {
		panic(err)
	}

	// 2. 创建自定义的 Transport 并设置 Proxy 属性
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 3. 构建自定义的 HTTP Client
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second, // 顺便设置超时时间
	}

	// 4. 发起请求测试
	resp, err := client.Get("https://www.cip.cc/")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("当前HTTP请求使用的IP:", string(body))
}
