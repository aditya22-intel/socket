package main

import (
	"os"
	"terminal/cmd"

	"github.com/aditya22-intel/socket/dcpcmd"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

var client *corev1client.CoreV1Client
var namespace string

func main() {
	arg := os.Args[1:]
	switch arg[0] {
	case "get":
		dcpcmd.GoGet(client, namespace, arg[1])
		break
	case "login":
		cmd.Login()
		client, namespace = cmd.GetOCClient()
		break
	default:
		panic("Please enter proper command")
	}

	//socket.SocketIO(client, namespace)
}
