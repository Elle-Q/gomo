package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gomo/config"
)

var (
	configYml      string
	ConfigStartCmd = &cobra.Command {
		Use:     "config",
		Short:   "Get Application config info",
		Example: "gomo config -c config/config.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

//初始化操作
func setup() {

	//1. 读取配置
	config.Setup("config/config.yml")
}

func init() {
	ConfigStartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yml", "Start server with provided configuration file")
}

func run() {
	//数据库配置
	database, errs := json.MarshalIndent(config.DatabaseConfig, "", "   ")
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println("database:", string(database))

	//app配置
	application, errs := json.MarshalIndent(config.ApplicationConfig, "", "   ")
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println("application:", string(application))
}

