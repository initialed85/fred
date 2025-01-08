package main

import (
	"log"

	"github.com/initialed85/fred/pkg/job_coordinator"
)

func main() {
	err := job_coordinator.Run()
	if err != nil {
		log.Fatal(err)
	}
}
