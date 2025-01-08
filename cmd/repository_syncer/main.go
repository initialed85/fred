package main

import (
	"log"

	"github.com/initialed85/fred/pkg/repository_syncer"
)

func main() {
	err := repository_syncer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
