#!/bin/bash

go build -o ./cmd/bin ./cmd/webapp-api

if [[ $1 = "run" ]]; then
  ./cmd/bin/webapp-api
fi
