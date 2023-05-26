package cmd

import (
	"context"
	"os"

	"go-odtec/cmd/yasuo"
	"go-odtec/cmd/yasuo/bootstrap"
	yasuoCfg "go-odtec/cmd/yasuo/config"
	"go-odtec/utils/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func makeRootCmd() {
	var (
		configPath string
		fileName   string
		fileType   string
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
		},
	}
	rootCmd.AddCommand(cmdYasuo)
}

func main() {
	makeRootCmd()
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
