package main

import (
	"main/client"
	"main/server"
	"time"
)

func main() {
	go server.Run()
	time.Sleep(time.Second)
	client.Run()
}
