/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/Xfers/cloudhsm-service-go/crypto"
	"github.com/libp2p/go-openssl"
	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var pureSignCmd = &cobra.Command{
	Use:   "pure-sign",
	Short: "Sign digest with a private key",
	Long: `Pure-sign is to sign an already digested value. 
	Usage is similar with regular sign. The only difference is it takes sha256 digest value.

	Sign from string (digested):
	$ go run main.go sign -k $YOUR_KEY_FILE -s "LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=%"
	
	Sign from file:
	$ go run main.go sign -k $YOUR_KEY_FILE -f $YOUR_FILE

	Sign from stdin:
	echo -n "hello" | openssl dgst -sha256 -binary - | base64 -w 0 | ./hsm-service pure-sign -k $YOUR_KEY_FILE 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read flags
		keyFile, _ := cmd.Flags().GetString("key")
		stringToSign, _ := cmd.Flags().GetString("string")
		fileToSign, _ := cmd.Flags().GetString("file")

		// Check flags
		if stringToSign != "" && fileToSign != "" {
			fmt.Println("Cannot sign from both string and file")
			return
		}

		// Check if key is specified
		if keyFile == "" {
			fmt.Println("Please specify key")
			return
		}

		var data string
		switch {
		case stringToSign != "":
			data = stringToSign
		case fileToSign != "":
			dataBa, err := ioutil.ReadFile(fileToSign)
			if err != nil {
				fmt.Println(err)
				return
			}
			data = string(dataBa)
		default:
			// Get data from stdin
			dataBa, err := ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				fmt.Println(err)
				return
			}
			data = string(dataBa)
			if data == "" {
				fmt.Println("Please provide data to sign")
				return
			}
		}

		// Key
		keyPem, err := crypto.GetKeyPem(&keyFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		key, err := openssl.LoadPrivateKeyFromPEM(keyPem)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Sign
		pureSigner := crypto.NewPureSigner(&key, data)
		signature, err := pureSigner.Sign()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(signature)
	},
}

func init() {
	rootCmd.AddCommand(pureSignCmd)

	// Here you will define your flags and configuration settings.

	// Key flag
	pureSignCmd.Flags().StringP("key", "k", "", "Private key file")
	err := pureSignCmd.MarkFlagRequired("key")
	if err != nil {
		fmt.Println(err)
		return
	}

	// String flag
	pureSignCmd.Flags().StringP("string", "s", "", "String to sign")

	// File flag
	pureSignCmd.Flags().StringP("file", "f", "", "File to sign")

	// // Stdin flag
	// pureSignCmd.Flags().BoolP(" ", " ", false, "Sign from stdin")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pureSignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pureSignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
