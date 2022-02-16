package cmd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

func GetPods(client *corev1client.CoreV1Client) {
	// List all Pods in our current Namespace.
	pods, err := coreclient.Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pods in namespace %s:\n", namespace)
	for _, pod := range pods.Items {
		fmt.Printf("  %s\n", pod.Name)
	}
}
