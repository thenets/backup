package main

import (
	"fmt"

	"github.com/thenets/backup/pkg/common"
)

func main() {
	fmt.Println("Hello World")

	// Import secret config file
	common.Loads("./secrets/key.yml")

	// Print priv and pub hash

	// Import SSH config file

	// Get distro info over SSH

	// Start sync
}

func checkRequirements() {
	// Check all binaries
}
