#!/bin/sh
# Exit if any command fails
set -e
cd /app
go mod download
go get ordersystem
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/ordersystem