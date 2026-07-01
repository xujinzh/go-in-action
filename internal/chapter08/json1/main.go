// 这个示例程序展示如何解码JSON字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON包含要反序列化的样例字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// 将JSON字符串反序列化到map变量
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Name:", c["name"])
	log.Println("Title:", c["title"])
	log.Println("Contact")
	log.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}
