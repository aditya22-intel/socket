package dcpcmd

import (
	"github.com/aditya22-intel/socket/cmd"
)

func GoGet(arg string) {

	switch arg {
	case "pods":
		cmd.GetPods()
		break
	default:
	}
}
