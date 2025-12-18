/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	
	"log"
	"strconv"
	"strings"

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/rawg"
	"github.com/spf13/cobra"
)


var askGameCmd = &cobra.Command{
	Use:   "askGame",
	Short: "ask for a game and see the results",
	Long: `ask for a game and see the results`,
	Run: func(cmd *cobra.Command, args []string) {
		game_name,err := cmd.Flags().GetString("gamename")
		if err!=nil{
			log.Fatal("provide a name")
		}
		com := rawg.Communicate{ApiKey: "076726af606e4c0d83c80068fc932955"}
		
		game, screenshotres := com.Ask(game_name)
		
		var builder strings.Builder
		builder.WriteString("Slug: ")
		builder.WriteString(*(game.Slug))
		builder.WriteString("\nDescription: ")
		builder.WriteString(*(game.Description))
		builder.WriteString("\nExternal ID: ")
		builder.WriteString(strconv.Itoa(*(game.ExternalID)))
		builder.WriteString("\nGenre: ")
		if game.Genres != nil && len(*game.Genres) > 0 {
			builder.WriteString((*game.Genres)[0].Slug)
		}
		builder.WriteString("\nPlatforms: ")
		
		builder.WriteString("\ndate: ")
		builder.WriteString(*game.Released)
		builder.WriteString("\nImage: ")
		builder.WriteString(*(game.Image))
		builder.WriteString("\nBackground Image: ")
		builder.WriteString(*(game.Main_image))
		builder.WriteString("\n")
		builder.WriteString(screenshotres.Results[0].Image)
		log.Println(builder.String())

	},
}

func init() {
	rootCmd.AddCommand(askGameCmd)

	askGameCmd.Flags().StringP("gamename", "g", "for-honor", "asks for game details")
}
