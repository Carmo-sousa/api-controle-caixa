package main

import (
	"github.com/Carmo-sousa/api/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}
