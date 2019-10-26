package main

import (
	"fmt"

	"github.com/thenets/backup/pkg/common"
)

func testSecret() {
	// Import secret config file
	secret := common.LoadSecret("./secrets/key.yml")

	// Print priv and pub hash
	fmt.Println(secret.GetPrivateKeyHash())
	fmt.Println(secret.GetPublicKeyHash())
}

func testSSH() {
	// Import SSH config file
	ssh := common.LoadSSH("./secrets/server.yml")

	// Get distro info over SSH
	// fmt.Printf("%#v", ssh)

	// Test connection
	

	// Start sync or dump
	ssh.Sync()

	// Compress

	// Status

	// Done

	fmt.Println("Synced")
}

func main() {
	fmt.Println("Hello World")

	// testSecret()
	testSSH()
}

func checkRequirements() {
	// Check all binaries
}
