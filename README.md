
# CloudHSM Service in Go

## How to generate test rsa key pairs

```bash
# Pair 1
openssl genrsa -out testrsaprivkey.pem 2048
openssl rsa -in testrsaprivkey.pem -outform PEM -pubout -out testpublic.pem

# Pair 2
openssl genrsa -out testrsaprivkey2.pem 2048
openssl rsa -in testrsaprivkey2.pem -outform PEM -pubout -out testpublic2.pem
```

## How to run the server

```bash
# run the server the signer services 
go run main.go --k1 testrsaprivkey.pem --k2 testrsaprivkey2.pem

# or like this (signer is the default mode):
go run main.go --mode signer--k1 testrsaprivkey.pem --k2 testrsaprivkey2.pem

# or shorter flag:
go run main.go -m signer --k1 testrsaprivkey.pem --k2 testrsaprivkey2.pem

# run the server with the verifier services
go run main.go --mode verifier --k1 testpublic.pem --k2 testpublic2.pem

 ```

 ## How to run the commands

```bash
# run the signer command

# Sign from string:
go run main.go sign -k $YOUR_KEY_FILE -s "hello"

# Sign from file:
go run main.go sign -k $YOUR_KEY_FILE -f $YOUR_FILE

# Sign from stdin:
cat $YOUR_FILE | go run main.go sign -k $YOUR_KEY_FILE
# or
echo "hello" | go run main.go sign -k $YOUR_KEY_FILE
```

```bash
# run the pure-sign command
echo -n "hello" | openssl dgst -sha256 -binary - | base64 -w 0 | go run main.go pure-sign -k $YOUR_KEY_FILE`,
```





