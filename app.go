package main

import (
	"fmt"
	"log"

	"github.com/thenets/backup/lib/config"
)

func main() {
	configFile, err := config.Loads("samples/minecraft-dir.yml")
	check(err)

	rsync, err := configFile.Rsync()
	check(err)
	fmt.Printf("%#v\n", rsync)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
