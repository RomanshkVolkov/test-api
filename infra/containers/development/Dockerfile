FROM golang:latest as builder
WORKDIR /srv

RUN go install github.com/air-verse/air@latest

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
