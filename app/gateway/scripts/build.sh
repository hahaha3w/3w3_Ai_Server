#!/bin/bash

# 设置环境变量
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# 编译 Go 程序
go build -o ../runner ../cmd/main.go

