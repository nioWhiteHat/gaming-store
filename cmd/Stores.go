/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/nioWhiteHat/gaming-store-backend.git/internal/rawg"
	"github.com/spf13/cobra"
)


var StoresCmd = &cobra.Command{
	Use:   "Stores",
	Short: "A brief description of your command",
	
	Run: func(cmd *cobra.Command, args []string) {
		var c rawg.Communicate
		c.ApiKey="076726af606e4c0d83c80068fc932955"
		ctx := context.Background()
		err := c.GetStores(ctx,dbpool)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println("Stores called")
	},
}

func init() {
	dbconCmd.AddCommand(StoresCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// StoresCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// StoresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
