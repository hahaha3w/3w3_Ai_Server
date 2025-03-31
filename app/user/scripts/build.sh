#!/bin/bash

# 设置环境变量
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# 编译 Go 程序
go build -o ../runner ../cmd/

GRANDPARENT_DIR_NAME=$(basename $(dirname $(pwd)))
IMAGE_NAME="xyq777/ai-biography-${GRANDPARENT_DIR_NAME}"
cd ..
# 构建 Docker 镜像
docker build  -t "$IMAGE_NAME" .
docker push "$IMAGE_NAME"
