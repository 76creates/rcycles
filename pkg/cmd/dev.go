package cmd

import (
	"github.com/rcycles/pkg/resource"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Short: "dev",
	Run: func(cmd *cobra.Command, args []string) {
		file := new(resource.File)
		file.Lookup()
	},
}