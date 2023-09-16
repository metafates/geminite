package main

import (
	"log"

	"github.com/metafates/geminite/cmd"
	"github.com/metafates/geminite/config"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
