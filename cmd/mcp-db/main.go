package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tamaco489/dbpilot/internal/config"
)

var cfgFile string

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dbpilot",
		Short: "MCP server for MySQL/PostgreSQL",
		RunE:  run,
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: config.yaml)")
	cmd.Flags().String("driver", "", "database driver: mysql or postgres")
	cmd.Flags().String("dsn", "", "data source name")
	cmd.Flags().Bool("read-only", false, "enable read-only mode")
	cmd.Flags().Int("max-open-conns", 10, "max open connections")
	cmd.Flags().Int("max-idle-conns", 5, "max idle connections")
	cmd.Flags().Int("conn-max-lifetime", 3600, "connection max lifetime (seconds)")

	_ = viper.BindPFlag("driver", cmd.Flags().Lookup("driver"))
	_ = viper.BindPFlag("dsn", cmd.Flags().Lookup("dsn"))
	_ = viper.BindPFlag("read_only", cmd.Flags().Lookup("read-only"))
	_ = viper.BindPFlag("max_open_conns", cmd.Flags().Lookup("max-open-conns"))
	_ = viper.BindPFlag("max_idle_conns", cmd.Flags().Lookup("max-idle-conns"))
	_ = viper.BindPFlag("conn_max_lifetime", cmd.Flags().Lookup("conn-max-lifetime"))

	cobra.OnInitialize(initConfig)

	return cmd
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	viper.SetEnvPrefix("DBPILOT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}

func run(cmd *cobra.Command, _ []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	cmd.Printf("driver: %s, read_only: %v\n", cfg.Driver, cfg.ReadOnly)
	// TODO: MCP サーバーの起動処理をここに追加
	return nil
}
