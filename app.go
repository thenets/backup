package main

import (
	"log"

	"github.com/thenets/backup/lib/config"
)

func main() {
	config.CheckConfigFileKind("samples/minecraft-dir.yml")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
