package main

import (
	"log"

	"github.com/initialed85/fred/pkg/trigger_producer"
)

func main() {
	err := trigger_producer.Run()
	if err != nil {
		log.Fatal(err)
	}
}
