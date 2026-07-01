package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

type (
	// gResult 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// gResponse 包含顶级的文档
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	// 1. 解析代理地址
	proxyURL, err := url.Parse("http://127.0.0.1:8118")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// 2. 创建自定义的 Transport 并设置 Proxy 属性
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 3. 构建自定义的 HTTP Client
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	// 4. 发起请求
	resp, err := client.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("API returned status %d. Response: \n%s", resp.StatusCode, string(body))
	}

	// 将 JSON 响应解码到结构类型
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}
