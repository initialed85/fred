package main

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	_, err := git.PlainClone(
		"/tmp/foo",
		false,
		&git.CloneOptions{
			URL:      "https://github.com/go-git/go-git",
			Progress: os.Stdout,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
