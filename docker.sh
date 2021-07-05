#!/bin/sh

echo "Building docker image..."

docker build . -t teodorstandavid/stock-checker:latest

# echo "Pushing Docker iamage to Docker hub"

# docker push teodorstandavid/stock-checker:latest