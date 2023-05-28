package cmd

import (
	"context"
	"os"

	"go-odtec/cmd/yasuo"
	"go-odtec/cmd/yasuo/bootstrap"
	yasuoCfg "go-odtec/cmd/yasuo/config"
	configs "go-odtec/configs/common"
	"go-odtec/utils/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func makeRootCmd() {
	var (
		configPath  string
		fileName    string
		fileType    string
		migratePath string
	)
	cmdYasuo := &cobra.Command{
		Use:   "yasuo",
		Short: "Start yasuo server",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &yasuoCfg.Config{}
			rsc := &bootstrap.Resources{}
			config.MustLoadConfig(
				configPath,
				fileName,
				fileType,
				cfg,
			)
			yasuo.StartGRPCServer(rsc, cfg)
		}}
	// 	cmdYone := &cobra.Command{
	// Use: "yone",
	// 	Short: "Start yone server",
	// 	Args: cobra.MinimumNArgs(0),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		cfg := &yasuoCfg.Config{}
	// 		rsc := &bootstrap.Resources{}
	// 		config.MustLoadConfig(
	// 			configPath,
	// 			fileName,
	// 			fileType,
	// 			cfg,
	// 		)
	// 		yone.StartGRPCServer(rsc, cfg)
	// 	}
	// }
	cmdMigrateSql := &cobra.Command{
		Use:   "sql_migrate",
		Short: "Run migration all services",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &configs.Config{}
			config.MustLoadConfig(
				configPath,
				fileName,
				fileType,
				cfg,
			)
			runDBMigration(cfg.MigrationURL, cfg.DBSource)
		}}
	cmdMigrateSql.PersistentFlags().StringVar(
		&migratePath,
		"migratePath",
		"file:///migrations",
		"path migration folder")
	rootCmd.PersistentFlags().StringVar(
		&configPath,
		"configPath",
		"",
		"path to common configuration file, usually used for configuration",
	)
	rootCmd.PersistentFlags().StringVar(
		&fileName,
		"fileName",
		"",
		"name of specific config file",
	)
	rootCmd.PersistentFlags().StringVar(
		&fileType,
		"fileType",
		"",
		"type of specific config file like yaml, yml, env, toml",
	)
	rootCmd.AddCommand(cmdYasuo)
	// rootCmd.AddCommand(cmdYone)
	rootCmd.AddCommand(cmdMigrateSql)

}

func main() {
	makeRootCmd()
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
