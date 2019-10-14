package main

import (
	"fmt"

	"github.com/thenets/backup/pkg/common"
)

func main() {
	fmt.Println("Hello World")

	// Import secret config file
	config := common.Loads("./secrets/key.yml")

	// Print priv and pub hash
	secret := config.SpecSecret()
	fmt.Println(secret.GetPrivateKeyHash())
	fmt.Println(secret.GetPublicKeyHash())
	// fmt.Println(secret.PrivateKey)
	// fmt.Println(secret.PublicKey)

	// Import SSH config file

	// Get distro info over SSH

	// Start sync
}

func checkRequirements() {
	// Check all binaries
}
