/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	
	"os"

	
	


	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "gaming-store-backend.git",
	Short: "A brief description of your application",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) { },
	
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
 

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gaming-store-backend.git.yaml)")

	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


