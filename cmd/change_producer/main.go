package main

import (
	"log"

	"github.com/initialed85/fred/pkg/change_producer"
)

func main() {
	err := change_producer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
