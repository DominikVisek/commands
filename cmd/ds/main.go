package main

import (
	"os"

	"ds/cmd/ds/cmd"
	"ds/pkg/output"
)

func main() {
	if os.Geteuid() == 0 && len(os.Args) > 1 && os.Args[1] != "hostname" {
		output.UserOut.Fatal("ddev is not designed to be run with root privileges, please run as normal user and without sudo")
	}

	cmd.Execute()
}
