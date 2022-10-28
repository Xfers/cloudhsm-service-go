
# CloudHSM Service in Go

## How to generate test rsa key pairs

```bash
openssl genrsa -out testrsaprivkey.pem 2048
openssl rsa -in testrsaprivkey.pem -outform PEM -pubout -out testpublic.pem
```

## How to run the server

```bash
go run main.go --private_key_path testrsaprivkey.pem --public_key_path testpublic.pem
```

