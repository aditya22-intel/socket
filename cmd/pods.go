package cmd

import (
	"fmt"

	"github.com/aditya22-intel/socket/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPods() {
	client, namespace := client.GetOCClient()
	//List all Pods in our current Namespace.
	pods, err := client.Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pods in namespace %s:\n", namespace)
	for _, pod := range pods.Items {
		fmt.Printf("  %s\n", pod.Name)
	}
}
