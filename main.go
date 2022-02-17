package main

import (
	"fmt"
	"os"

	"github.com/aditya22-intel/socket/cmd"

	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

var client corev1client.CoreV1Client
var namespace string

func main() {
	arg := os.Args[1:]
	switch arg[0] {
	case "get":
		fmt.Println("new one")
		cmd.GetPods(client, namespace)
		//dcpcmd.GoGet(client, namespace, arg[1])
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
