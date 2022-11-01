/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pureSignCmd represents the pureSign command
var pureSignCmd = &cobra.Command{
	Use:   "pure-sign",
	Short: "Sign digest with a private key",
	Long: `Pure-sign is to sign an already digested value. 
	Usage is similar with regular sign. The only difference is it takes sha256 digest value.
	For example:
	
	echo -n "hello" | openssl dgst -sha256 -binary - | base64 -w 0 | go run main.go pure-sign -k $YOUR_KEY_FILE`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pureSign called")
	},
}

func init() {
	rootCmd.AddCommand(pureSignCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pureSignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pureSignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
