package main

import (
	"terminal/cmd"
	"terminal/socket"
)

func main() {
	cmd.Login()
	client, namespace := cmd.GetOCClient()
	socket.SocketIO(client, namespace)
}
