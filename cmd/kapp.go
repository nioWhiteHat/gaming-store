package cmd

import (
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/config"
	"github.com/spf13/cobra"
)

var forceCmd = &cobra.Command{
	Use:   "force [version]",
	Short: "Forces the database to a specific version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid version number provided: %v", err)
		}

		cfg := config.LoadDBConfig()
		m, err := migrate.New("file://migrations", cfg.DSNForMigrate)
		if err != nil {
			log.Fatalf("Failed to create migrate instance: %v", err)
		}

		log.Printf("Forcing database to version %d...", version)
		if err := m.Force(version); err != nil {
			log.Fatalf("Failed to force migration: %v", err)
		}
		log.Printf("Successfully forced database to version %d.", version)
	},
}

func init() {
	migrateCmd.AddCommand(forceCmd)
}
