#!/bin/bash
GRANDPARENT_DIR_NAME=$(basename $(dirname $(pwd)))
IMAGE_NAME="xyq777/ai-biography-${GRANDPARENT_DIR_NAME}"
cd ..
# 构建 Docker 镜像
docker build -t $IMAGE_NAME .
docker push $IMAGE_NAME
