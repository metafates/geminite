package cmd

import (
	"context"
	"net/url"

	"github.com/metafates/geminite/page"
	"github.com/metafates/geminite/tui"
	"github.com/metafates/geminite/tui/state/pageview"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geminite [URL]",
	Short: "Article reader for your terminal",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		URL, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		p, err := page.New(context.Background(), URL)
		if err != nil {
			return err
		}

		if err := tui.Run(pageview.New(p)); err != nil {
			return err
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
