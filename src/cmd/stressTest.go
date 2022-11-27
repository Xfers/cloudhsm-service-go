/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Xfers/cloudhsm-service-go/api"
	"github.com/spf13/cobra"
)

// stressTestCmd represents the stressTest command
var stressTestCmd = &cobra.Command{
	Use:   "stressTest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.StressTest()
	},
}

func init() {
	rootCmd.AddCommand(stressTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stressTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stressTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
