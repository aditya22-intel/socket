package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aditya22-intel/socket/utils"
)

func Login() {
	utils.DownloadAndUncompress("oc.gz")
	fmt.Print("Enter Openshift Username: ")
	var user string
	var password string
	fmt.Scanln(&user)

	fmt.Print("Enter Openshift Password: ")
	fmt.Scanln(&password)

	fmt.Println()
	//cmd := exec.Command("bash", "-c", "echo "+user+" "+string(password)+"| oc login")
	cmd := exec.Command("./oc", "login", "https://api.devcloud.awsdevcloud.com:6443", "-u", user, "-p", password)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Remove("oc")
}
