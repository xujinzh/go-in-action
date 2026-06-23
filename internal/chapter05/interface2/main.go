package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	// 将字符串写入 Buffer
	b.Write([]byte("Hello "))
	// 使用 Fprintf 将字符串拼接到 Buffer
	fmt.Fprintf(&b, "World!")
	// 将 Buffer 的内容写到 Stdout
	io.Copy(os.Stdout, &b)
}
