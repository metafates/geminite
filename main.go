package main

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/tui"
	"github.com/metafates/geminite/tui/state/pageview"
)

func main() {
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	p, err := page.New(context.Background(), URL)
	if err != nil {
		log.Fatal(err)
	}

	if err := tui.Run(pageview.New(p)); err != nil {
		log.Fatal(err)
	}
}
