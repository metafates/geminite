package cmd

import (
	"context"
	"net/url"

	"github.com/metafates/geminite/browser"
	"github.com/metafates/geminite/tui"
	"github.com/metafates/geminite/tui/state/bookmarkview"
	"github.com/metafates/geminite/tui/state/pageview"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geminite [URL]",
	Short: "Article reader for your terminal",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		browser, err := browser.New()
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return tui.Run(bookmarkview.New(browser))
		}

		URL, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		p, err := browser.Open(context.Background(), URL)
		if err != nil {
			return err
		}

		if err := tui.Run(pageview.New(browser, p)); err != nil {
			return err
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
