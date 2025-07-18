package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	// так нужно.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// так нужно.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Выполнить миграцию базы данных",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("migration run...")

		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBRepo.User,
			cfg.DBRepo.Password,
			cfg.DBRepo.Host,
			cfg.DBRepo.Port,
			cfg.DBRepo.DB,
		)

		m, err := migrate.New("file://migrations", dsn)
		if err != nil {
			log.Fatalf("migrator creating: %v", err)
		}

		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("mirgate: %v", err)
		}

		fmt.Println("success migration")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
