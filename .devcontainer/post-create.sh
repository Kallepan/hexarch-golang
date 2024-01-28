#!/bin/bash

# install go tools
go install github.com/google/wire/cmd/wire@latest \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && go install github.com/bufbuild/buf/cmd/buf@v1.28.1

# install go modules
cd src \
    && go mod tidy

echo "Done!"