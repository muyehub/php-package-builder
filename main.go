package main

import (
	"log"

	"github.com/muyehub/php-package-builder/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
