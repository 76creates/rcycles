package cmd

import (
	"github.com/rcycles/pkg/config"
	"log"

	"cd ../pkg/tools"
	"github.com/spf13/cobra"
)

var (
	version = "0"
	configPath string
	cfg *config.RCyclesConfig
)

func Main() {
	rootCmd.PersistentFlags().StringVarP(
		&configPath, "config", "c", "./config.yaml",
		"relative or absolute path to the config file",
	)

	rootCmd.AddCommand(devCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "rcycles",
	Short: "service for managing garbage on resources",
	Long: tools.DeIndent(`
	RCycles service is used for seeking and reporting on unused resources based
	on plugins added. For now AWS and Vayu are two main sources of infromation
	regarding whats ready for purging, for both sources there are shared resources
	that exist to minimize number of calls towards AWS mainly, and VayuDB.
	`),
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
