/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/rawg"
	"github.com/spf13/cobra"
)

// CreatePlatformsCmd represents the CreatePlatforms command
var CreatePlatformsCmd = &cobra.Command{
	Use:   "CreatePlatforms",
	Short: "A brief description of your command",
	
	Run: func(cmd *cobra.Command, args []string) {

		com := rawg.Communicate{ApiKey: "076726af606e4c0d83c80068fc932955"}
		var ctx context.Context
		com.GetPlatforms(ctx,dbpool)
	},
}

func init() {
	dbconCmd.AddCommand(CreatePlatformsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CreatePlatformsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CreatePlatformsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
