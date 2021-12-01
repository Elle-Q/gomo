package cmd

import (
	"errors"
	"fmt"
	"gomo/cmd/config"
	"gomo/tool"
	"os"

	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:          "gomo-admin",
	Short:        "gomo-admin",
	SilenceUsage: true,
	Long:         `gomo-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(tool.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + tool.Green(`gomo-admin `) + ` 可以使用 ` + tool.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	//rootCmd.AddCommand(api.StartCmd)
	//rootCmd.AddCommand(migrate.StartCmd)
	//rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(config.StartCmd)  //config加载闭包
	//rootCmd.AddCommand(app.StartCmd)
}


//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}