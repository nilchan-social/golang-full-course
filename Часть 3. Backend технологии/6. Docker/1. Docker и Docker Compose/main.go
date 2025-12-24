package main

import (
	"fmt"
	"study/http_server"
)

func main() {
	if err := http_server.StartHTTPServer(); err != nil {
		fmt.Println("HTTP server stopped:", err)
	}
}
