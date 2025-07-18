package commands

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/infrastructure/app"
	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/infrastructure/repository/bucketrepo"
	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/infrastructure/repository/configrepo"
	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/infrastructure/server"
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
		log.Println("Starting anti-bruteforce...")
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
	configRepo, err := configrepo.NewPgConfigRepo(cfg.DBRepo)
	if err != nil {
		return fmt.Errorf("create configRepo repository: %w", err)
	}

	bucketRepo, err := bucketrepo.NewRedisBucketRepo(cfg.MemoryRepo)
	if err != nil {
		return fmt.Errorf("create bucketRepo repository: %w", err)
	}

	antiBruteForceApp := app.NewAntiBruteForceApp(configRepo, bucketRepo)
	serverGrpc := server.NewGrpcServer(cfg.GrpcServer, *antiBruteForceApp)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	log.Println("Starting server...")

	go func() {
		if err := serverGrpc.Start(ctx); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serverGrpc.Stop(shutdownCtx); err != nil {
		log.Printf("gRPC server shutdown error: %v", err)
	} else {

		log.Println("gRPC server shutdown complete")
	}

	return nil
}
