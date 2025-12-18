/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/data"
	"github.com/nioWhiteHat/gaming-store-backend.git/internal/rawg"
	"github.com/spf13/cobra"
)


var InsertGame = &cobra.Command{
	Use:   "InsertGame",
	Short: "asks for details of a game in the rawg api",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		game_name,err := cmd.Flags().GetString("gamename1")
		if err!=nil{
			log.Fatal("provide a name")
		}
		com := rawg.Communicate{ApiKey: "076726af606e4c0d83c80068fc932955"}
		var ctx context.Context
		game, screenshotres := com.Ask(game_name)
		gamemodel := data.GameModel{DB: dbpool}
		gamemodel.InsertGame(ctx, game, screenshotres)
		
		 
		
	},
}

func init() {
	dbconCmd.AddCommand(InsertGame)

	InsertGame.Flags().StringP("gamename1", "s", "for-honor", "asks for game details")
}
