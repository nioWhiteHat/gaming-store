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

var dbpool *pgxpool.Pool

var dbconCmd = &cobra.Command{
	Use:   "dbcon",
	Short: "connects the children cmds with the database",
	Long:  ``,
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
		dbpool = pool
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dbcon called")
	},
}

func init() {
	rootCmd.AddCommand(dbconCmd)

}
