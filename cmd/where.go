package cmd

import (
	"fmt"

	"github.com/metafates/geminite/where"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(whereCmd)
}

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Available paths",
	Args:  cobra.NoArgs,
	Run: func(*cobra.Command, []string) {
		type Target struct {
			Name string
			Path string
		}

		for _, target := range []Target{
			{
				Name: "Config file",
				Path: where.ConfigFile(),
			},
		} {
			fmt.Println(target.Name)
			fmt.Println(target.Path)
			fmt.Println()
		}
	},
}
