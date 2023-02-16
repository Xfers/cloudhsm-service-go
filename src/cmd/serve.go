/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"strconv"

	"github.com/Xfers/cloudhsm-service-go/api"
	"github.com/spf13/cobra"
)

// Define const map with iota for signer and verifier modes
const (
	Signer = iota
	Verifier
)

// Define map for modes
var modeMap = map[string]int{
	"signer":   Signer,
	"verifier": Verifier,
}

// Define default mode
const defaultMode = "signer"

// Define maximum arguments for keys
const maxKeys = 10

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server",
	Long: `Run the server to serve the API. 

	Example to run signer services digest, sign, and pure-sign: 

	go run main.go serve -m signer -k1 key.pem -k2 key2.pem -k3 key3.pem

	Example to run verifier service:

	go run main.go serve -m verifier -k1 key.pem -k2 key2.pem -k3 key3.pem`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get mode from flag
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			panic(err)
		}
		_, ok := modeMap[mode]
		if !ok {
			panic(errors.New("Invalid mode"))
		}

		// Get keys from flags
		keys := map[string]string{}
		for i := 1; i <= maxKeys; i++ {
			key, err := cmd.Flags().GetString("k" + strconv.Itoa(i))
			if err != nil {
				panic(err)
			}
			if key != "" {
				keys["k"+strconv.Itoa(i)] = key
			}
		}

		// Check if keys set
		if len(keys) == 0 {
			panic(errors.New("No keys set"))
		}

		// Run Signer server
		if mode == defaultMode {
			api.RunSignerServer(map[string]interface{}{
				"keys": keys,
			})
		} else {
			api.RunVerifierServer(map[string]interface{}{
				"keys": keys,
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Add flag called mode can be signer which is default and verifier
	serveCmd.Flags().StringP("mode", "m", defaultMode, "Mode of the server, can be signer or verifier")

	// Add flag called keys liek this k1, k2, maximum 10 keys
	for i := 1; i <= maxKeys; i++ {
		name := "k" + strconv.Itoa(i)
		serveCmd.Flags().StringP(name, "", "", "Key "+strconv.Itoa(i))
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
