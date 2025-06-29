package main

import (
	"os"
	"path/filepath"

	"github.com/celer-network/intent-rfq-relayer/relayer/service/relayer"

	"github.com/celer-network/goutils/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagHome     = "home"
	FlagLoglevel = "loglevel"
	FlagLogDir   = "logdir"
	Port         = "mm.port"
)

func main() {
	rootCmd := NewRootCmd()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "intent-rfq-relayer",
		Short: "rfq relayer",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			home, err := cmd.Flags().GetString(FlagHome)
			checkErr(err)
			viper.SetDefault(FlagHome, home)

			ll, err := cmd.Flags().GetString(FlagLoglevel)
			checkErr(err)
			log.SetLevelByName(ll)

			ld, err := cmd.Flags().GetString(FlagLogDir)
			checkErr(err)
			log.SetDir(ld)

			setupConfig(home)
			return nil
		},
	}
	rootCmd.PersistentFlags().String(FlagHome, os.ExpandEnv("$HOME/.intent-rfq-relayer"), "home path")
	rootCmd.PersistentFlags().String(FlagLoglevel, "info", "log level")
	rootCmd.PersistentFlags().String(FlagLogDir, "$HOME/.intent-rfq-relayer/app", "log level")
	rootCmd.AddCommand(
		startCmd(),
	)
	return rootCmd
}

func setupConfig(home string) {
	log.Infoln("Reading example-mm configs")

	readConfig(home, "config/chain.toml")
	readConfig(home, "config/fee.toml")
	readConfig(home, "config/lp.toml")
	readConfig(home, "config/mm.toml")
	readConfig(home, "config/relayer.toml")
}

func readConfig(home, relativePath string) {
	path := filepath.Join(home, relativePath)
	viper.SetConfigFile(path)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalln("failed to load", path, err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("fatal error occurred: %s", err.Error())
	}
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start relayer",
		Run: func(cmd *cobra.Command, args []string) {
			relayerServer := relayer.NewRfqRelayerServer()
			relayerServer.ReportConfigs()
			relayerServer.DefaultProcessOrder()
			relayerServer.Serve()
		},
	}
	return cmd
}
