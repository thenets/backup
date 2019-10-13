package main

import (
	"fmt"

	"github.com/thenets/backup/pkg/ssh"
)

func main() {
	fmt.Println("Hello World")
	ssh.RunAll("/tmp/dirBkp.yaml")
}

func checkRequirements() {
	// Check all binaries
}
