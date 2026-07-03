package main

import (
	"encoding/json"
	"net/http"
)

// Routes 为网络服务设置路由
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON 返回一个简单的JSON文档
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@outlook.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
