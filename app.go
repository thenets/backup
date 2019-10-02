package main

import (
	"fmt"

	"github.com/thenets/backup/connectors"
)

func main() {
	fmt.Println("Hello World")
	connectors.SSHRunAll("/tmp/dirBkp.yaml")
}

func checkRequirements() {
	// Check all binaries
}
