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
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign plain text with a private key",
	Long: `Sign plain text data with a private key . For example:

	Sign from string:
	$ go run main.go sign -k $YOUR_KEY_FILE -s "hello"
	
	Sign from file:
	$ go run main.go sign -k $YOUR_KEY_FILE -f $YOUR_FILE

	Sign from stdin:
	$ cat $YOUR_FILE | go run main.go sign -k $YOUR_KEY_FILE
	or
	$ echo "hello" | go run main.go sign -k $YOUR_KEY_FILE 
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

		// Sign
		signer := crypto.NewSigner(&key, data)
		signature, err := signer.Sign()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(signature)
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Key flag
	signCmd.Flags().StringP("key", "k", "", "Private key file")
	err := signCmd.MarkFlagRequired("key")
	if err != nil {
		fmt.Println(err)
		return
	}

	// String flag
	signCmd.Flags().StringP("string", "s", "", "String to sign")

	// File flag
	signCmd.Flags().StringP("file", "f", "", "File to sign")

	// Stdin flag
	signCmd.Flags().BoolP("stdin", "", false, "Sign from stdin")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
