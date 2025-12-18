/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/rawg"
	"github.com/spf13/cobra"
)

// fillGenresCmd represents the fillGenres command
var fillGenresCmd = &cobra.Command{
	Use:   "fillGenres",
	Short: "it fills all the genres in the table",
	Long: `it fills all the genres in the table`,
	Run: func(cmd *cobra.Command, args []string) {
		com := rawg.Communicate{ApiKey: "076726af606e4c0d83c80068fc932955"}
		
		ctx := context.Background()
		com.GetGenres(dbpool, ctx)
	},
}

func init() {
	dbconCmd.AddCommand(fillGenresCmd)

	
}
