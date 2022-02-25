package cmd

import (
	"fmt"
	"os"

	"github.com/aditya22-intel/socket/client"
	"github.com/openshift/oc/pkg/cli/rsh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func GetPods() {
	client, namespace, _ := client.GetOCClient()
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

func GetShell() {
	_, namespace, restclient := client.GetOCClient()
	in, out, err := os.Stdin, os.Stdout, os.Stderr
	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: err}
	obj := rsh.NewRshOptions(ioStreams)
	obj.ExecOptions.Namespace = namespace
	obj.ExecOptions.Config = restclient
	obj.ExecOptions.StreamOptions.ContainerName = "result-service-devcloud-dev"
	obj.ExecOptions.StreamOptions.PodName = "result-service-devcloud-dev-17-rd4dw"
	obj.ExecOptions.StreamOptions.TTY = true
	obj.Executable = "/bin/sh"
	obj.ExecOptions.Command = []string{"/bin/sh"}
	obj.ForceTTY = false
	obj.Run()
}
