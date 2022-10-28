package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/Xfers/cloudhsm-service-go/crypto"
)

func main() {

	// Set test keys and data
	privatekey_pem, publickey_pem := generateTestKeys()
	data := "test data that is to be signed"

	// Digest
	d, err := crypto.Digest(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Digest: ", d)
	fmt.Println("")
	//print line
	printLine()

	// Set Signers
	signer := crypto.NewSigner(privatekey_pem, d)                   //default, using go's package
	signerOpenSSL := crypto.NewSigner(privatekey_pem, d, "openssl") //openssl

	// Sign
	signature, err := signer.Sign()
	if err != nil {
		panic(err)
	}
	fmt.Println("Signature: ", signature)
	fmt.Println("")

	signatureOpenSSl, err := signerOpenSSL.Sign()
	if err != nil {
		panic(err)
	}
	fmt.Println("Signature OpenSSL: ", signatureOpenSSl)
	fmt.Println("")

	printLine()

	// Test Verify correct public key
	verifier := crypto.NewVerifier(publickey_pem, signature, data)
	verifierOpenSSL := crypto.NewVerifier(publickey_pem, signatureOpenSSl, data, "openssl")

	msg := "Verify correct public key: "

	result := verifier.Verify()
	printResult(msg, result)
	fmt.Println("")

	result = verifierOpenSSL.Verify()
	printResult("OpenSSL "+msg, result)
	fmt.Println("")

	printLine()

	// Test Verify wrong public key
	wrongPubKeyVerifier := crypto.NewVerifier("wrong public key", signature, data)
	wrongPubKeyVerifierOpenSSL := crypto.NewVerifier("wrong public key", signatureOpenSSl, data, "openssl")

	msg = "Result for wrong public key:"

	result = wrongPubKeyVerifier.Verify()
	printResult(msg, result)
	fmt.Println("")

	result = wrongPubKeyVerifierOpenSSL.Verify()
	printResult("OpenSSL "+msg, result)
	fmt.Println("")

	printLine()

	// Test Verify wrong data
	wrongDataVerifier := crypto.NewVerifier(publickey_pem, signature, "wrong data")
	wrongDataVerifierOpenSSL := crypto.NewVerifier(publickey_pem, signatureOpenSSl, "wrong data", "openssl")

	msg = "Result with wrong data:"

	result = wrongDataVerifier.Verify()
	printResult(msg, result)
	fmt.Println("")

	result = wrongDataVerifierOpenSSL.Verify()
	printResult("OpenSSL "+msg, result)
	fmt.Println("")

	printLine()

	// Test Verify wrong signature
	wrongSignatureVerifier := crypto.NewVerifier(publickey_pem, "wrong signature", data)
	wrongSignatureVerifierOpenSSL := crypto.NewVerifier(publickey_pem, "wrong signature", data, "openssl")

	msg = "Result with wrong signature:"

	result = wrongSignatureVerifier.Verify()
	printResult(msg, result)
	fmt.Println("")

	result = wrongSignatureVerifierOpenSSL.Verify()
	printResult("OpenSSL "+msg, result)
	fmt.Println("")

	printLine()
}

func generateTestKeys() (string, string) {
	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Cannot generate RSA key: %v", err)
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey

	// dump private key to string
	privatekey_pem := string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey),
	}))

	// dump public key to string
	publickey_pem := string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publickey),
	}))

	//print the keys
	fmt.Println("Private Key: ", privatekey_pem)
	fmt.Println("Public Key: ", publickey_pem)
	printLine()

	return privatekey_pem, publickey_pem
}

func printLine() {
	fmt.Println("------------------------------------------------------------")
}

func printResult(msg string, result bool) {
	if result {
		fmt.Println(msg, "Result: Success")
	} else {
		fmt.Println(msg, "Result: Fail")
	}
}
