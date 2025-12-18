/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/config"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/database"
	"github.com/spf13/cobra"
)
var dbpoolM *pgxpool.Pool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadDBConfig()
		log.Printf("DEBUG: The DSN string my program sees is: [%s]", cfg.DSNForApp)
		pool, err := database.NewDBPool(cfg.DSNForApp)
		if err != nil {
			log.Fatalf("Failed to create database connection pool: %v", err)
		}
		if err := pool.Ping(context.Background()); err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}
		log.Print("all good")
		dbpoolM = pool
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
