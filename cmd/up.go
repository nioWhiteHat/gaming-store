package cmd

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/config"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "creates all the tables",
	Long:  `creates all the tables`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadDBConfig()
		m, err := migrate.New(
			"file://migrations",
			cfg.DSNForMigrate,
		)
		if err != nil {
			log.Fatalf("failed to create migrate instance: %v", err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to apply migrations: %v", err)
		}
		
		
	},
}

func init() {
	migrateCmd.AddCommand(upCmd)

}
