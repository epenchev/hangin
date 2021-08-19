# base build image
FROM golang:1.16-buster

WORKDIR /ngha-build

# only copy go.mod and go.sum
COPY go.mod .
COPY go.sum .

RUN go mod download

# copy all the sources
COPY . /ngha-build
