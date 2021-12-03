package config

import "github.com/spf13/cobra"

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "config",
		Short:   "Get Application config info",
		Example: "gomo-admin config -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yml", "Start server with provided configuration file")
}

func run() {
	config.set
}

