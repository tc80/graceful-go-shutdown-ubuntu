FROM debian:stretch

WORKDIR /usr/go/src
ENV GOPATH=/usr/go

# avoid selecting location for tzdata
ARG DEBIAN_FRONTEND=noninteractive

RUN apt update
RUN apt-get install -y golang-go
