package main

import (
	"log"

	"github.com/metafates/geminite/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
