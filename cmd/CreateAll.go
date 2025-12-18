/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/rawg"
	"github.com/spf13/cobra"
)

// CreateAllCmd represents the CreateAll command
var CreateAllCmd = &cobra.Command{
	Use:   "CreateAll",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		
		com := rawg.Communicate{ApiKey: "076726af606e4c0d83c80068fc932955"}
		err := com.FillAllPipeline(dbpool,ctx)
		if err!=nil{
			log.Panicln(err)
		}
	},
}

func init() {
	dbconCmd.AddCommand(CreateAllCmd)

	
}
