package main

import (
	"time"

	"exc8/client"
	"exc8/server"
)

func main() {
	go func() {
		// todo start server
		server.StartGrpcServer()

	}()
	time.Sleep(1 * time.Second)
	// todo start client
	c, _ := client.NewGrpcClient()
	c.Run()
	println("Orders complete!")
}
