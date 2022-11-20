FROM golang:1.18-bullseye

RUN apt update && \
    apt install -y libssl-dev build-essential curl

RUN cd /tmp \
  && curl -LO -s https://s3.amazonaws.com/cloudhsmv2-software/CloudHsmClient/Bionic/cloudhsm-dyn_latest_u18.04_amd64.deb \
  && apt install -y ./cloudhsm-dyn_latest_u18.04_amd64.deb \
  && rm ./cloudhsm-dyn_latest_u18.04_amd64.deb

WORKDIR /go/src/hsm-service
COPY src/ .

RUN go mod download
RUN go mod verify

RUN go build -o hsm-service .

EXPOSE 8000
ENTRYPOINT ["./hsm-service", "serve"]