#!/bin/sh

export SYMBOL="MSFT"
export API_KEY="M4BNP375Y9LGWZQO"
export NDAYS=7
export GIN_MODE="debug"

go mod tidy

go run main.go
