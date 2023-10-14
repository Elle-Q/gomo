package cmd

import (
	"leetroll/cmd/config"
	"leetroll/cmd/server"
	"os"

	//"errors"
	"fmt"
	"github.com/spf13/cobra"
	"leetroll/tool"
)

var rootCmd = &cobra.Command{
	Use:          "gomo",
	Short:        "gomo",
	SilenceUsage: true,
	Long:         `gomo`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `welcome use >>> ` + tool.Green(`gomo-admin `)
	fmt.Printf("%s\n", usageStr)

}

func init() {
	//rootCmd.AddCommand(api.ConfigStartCmd)
	//rootCmd.AddCommand(migrate.ConfigStartCmd)
	rootCmd.AddCommand(config.ConfigStartCmd) //config加载闭包
	rootCmd.AddCommand(server.ServerStartCmd)
	//rootCmd.AddCommand(app.ConfigStartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
