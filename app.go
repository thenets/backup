package main

import (
	"fmt"

	"github.com/thenets/backup/pkg/common"
)

func main() {
	fmt.Println("Hello World")

	// Import secret config file
	secretConfig := common.Loads("./secrets/key.yml")

	// Print priv and pub hash
	secret := secretConfig.KindSecret()
	fmt.Println(secret.GetPrivateKeyHash())
	fmt.Println(secret.GetPublicKeyHash())

	// Import SSH config file
	sshConfig := common.Loads("./secrets/server.yml")

	// Get distro info over SSH
	fmt.Println(sshConfig.SSH())

	// Start sync
}

func checkRequirements() {
	// Check all binaries
}
