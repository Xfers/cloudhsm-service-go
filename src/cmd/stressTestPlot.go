/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"

	test "github.com/Xfers/cloudhsm-service-go/test/loadTest"
	"github.com/spf13/cobra"
)

// stressTestPlotCmd represents the stressTestPlot command
var stressTestPlotCmd = &cobra.Command{
	Use:   "stressTestPlot",
	Short: "load testing the cloudhsm service",
	Long: `Custom Load testing the cloudhsm service
	usage example: ./cloudhsm stressTestPlot -c 100 -e http://localhost:8080 -d hello`,
	Run: func(cmd *cobra.Command, args []string) {

		// read config file
		config, err := ioutil.ReadFile(cmd.Flag("config").Value.String())
		if err != nil {
			fmt.Println(err)
		}

		result, err := test.RunLoadTest(string(config))
		if err != nil {
			fmt.Println(err)
		}
		test.PlotterLoadTest(result)
	},
}

func init() {
	rootCmd.AddCommand(stressTestPlotCmd)

	// Here you will define your flags and configuration settings.
	// add flag for loading json config
	stressTestPlotCmd.Flags().StringP("config", "c", "", "config file path")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stressTestPlotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stressTestPlotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
