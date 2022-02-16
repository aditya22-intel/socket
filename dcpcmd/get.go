package dcpcmd

import (
	"github.com/aditya22-intel/socket/cmd"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

func GoGet(client *corev1client.CoreV1Client, namespace string, arg string) {

	switch arg {
	case "pods":
		cmd.GetPods(client, namespace)
		break
	default:
	}
}
