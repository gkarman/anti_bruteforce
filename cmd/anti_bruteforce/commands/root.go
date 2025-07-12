package commands

import (
	"fmt"
	"log"

	"github.com/gkarman/anti_bruteforce/internal/config"
	"github.com/spf13/cobra"
)

var (
	configPath string
	cfg        *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "anti_bruteforce",
	Short: "Start anti_bruteforce application",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		var err error
		cfg, err = config.Load(configPath)
		if err != nil {
			return fmt.Errorf("process load config: %w", err)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting anti-bruteforce...")
		if err := runApp(); err != nil {
			log.Fatalf("Fail start: %v", err)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "configs/anti_bruteforce_config.yaml", "configuration")
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Fail start: %v", err)
	}
}

func runApp() error {
	fmt.Println(cfg.GrpcServer)
	fmt.Println(cfg.InMemoryRepository)
	fmt.Println(cfg.SQLRepository)
	return nil
}
