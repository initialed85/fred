package main

import (
	"log"

	"github.com/initialed85/fred/pkg/job_executor"
)

func main() {
	err := job_executor.Run()
	if err != nil {
		log.Fatal(err)
	}
}
