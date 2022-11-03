
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

## How To Build the Executable

```bash
# build the executable
go build -o hsm-service main.go 
```

## How to run the server

```bash
# run the server the signer services 
./hsm-service serve --k1 testrsaprivkey.pem --k2 testrsaprivkey2.pem

# or like this (signer is the default mode, so that is why it was omitted above ):
./hsm-service serve --mode signer --k1 testrsaprivkey.pem --k2 testrsaprivkey2.pem

# or shorter flag (-m):
./hsm-service serve -m signer --k1 testrsaprivkey.pem --k2 testrsaprivkey2.pem

# run the server with the verifier services
./hsm-service serve --mode verifier --k1 testpublic.pem --k2 testpublic2.pem

 ```

 ## How to run the commands

```bash
# run the sign command

# Sign from string:
./hsm-service sign -k $YOUR_KEY_FILE -s "hello"

# Sign from file:
./hsm-service sign -k $YOUR_KEY_FILE -f $YOUR_FILE

# Sign from stdin:
cat $YOUR_FILE | ./hsm-service sign -k $YOUR_KEY_FILE
# or
echo -n "hello" | ./hsm-service sign -k $YOUR_KEY_FILE
```

```bash
# run the pure-sign command

# Sign from string(digested "hello"):
./hsm-service sign -k $YOUR_KEY_FILE -s "LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ="

# Sign from file (that holds the digest):
./hsm-service sign -k $YOUR_KEY_FILE -f $YOUR_DIGEST_FILE

# Sign from stdin:
echo -n "hello" | openssl dgst -sha256 -binary - | base64 -w 0 | ./hsm-service pure-sign -k $YOUR_KEY_FILE
```

## How to build the swagger docs

```bash
# build the swagger docs (swaggo is required: go install github.com/swaggo/swag/cmd/swag@latest )
swag init -g main.go --output docs
```

### Swagger docs are available at:

```url
 {HOST}:{PORT}/swagger/index.html
```
