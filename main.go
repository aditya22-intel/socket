package main

import (
	"github.com/aditya22-intel/socket/cmd"
	"github.com/aditya22-intel/socket/socket"
)

func main() {
	cmd.Login()
	client, namespace := cmd.GetOCClient()
	socket.SocketIO(client, namespace)
}
