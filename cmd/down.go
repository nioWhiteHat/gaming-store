/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/config"

	"github.com/spf13/cobra"
)


var downCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cfg:=config.LoadDBConfig()
		m,err := migrate.New(
			"file://migrations",
			cfg.DSNForMigrate)
		if err!=nil{
			log.Fatalf("couldnt make migrator, %s",err)
		}
		drop,er:= cmd.Flags().GetBool("all")
		if er!=nil{
			log.Fatal("u need to specify")
		}
		if drop{
			err = m.Drop()
			if err!=nil{
				log.Fatal(err)
			}
			log.Println("tables droped")
		}else{
			err = m.Down()
			if err!=nil{
				log.Fatal(err)
			}
		}
	},
}

func init() {
	migrateCmd.AddCommand(downCmd) 
	downCmd.Flags().Bool("all", false, "Apply all down migrations")
	
}
