FROM golang:1.18-bullseye

RUN apt update && \
    apt install -y libssl-dev build-essential

WORKDIR /go/src/hsm-service
COPY src/ .

RUN go mod download
RUN go mod verify

RUN go build -o hsm-service main.go

EXPOSE 8000
ENTRYPOINT ["./hsm-service", "serve"]