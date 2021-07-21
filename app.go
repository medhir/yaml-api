package main

import "github.com/medhir/yaml-api/server"

func main() {
	server := server.NewServer(":1111")
	server.Start()
}
