#!/bin/bash

# 该脚本可用于自动打包，上传到服务器，并重启指定服务
# 执行 chmod +x build.sh
# ./build.sh web_user

set -e

REMOTE_HOST="124.223.70.243"
TAG=1.0.0.0
NUMBER=6001
NAME="localhost:${NUMBER}/web_user"

build_service() {
  echo "name: $1"
  echo "go mod vendor.........."
  go mod vendor

  echo "begin to build ${NAME}:${TAG}..........."
  docker build -t ${NAME}:${TAG} .

  echo "begin to save "$1".tar"
  docker save -o "$1".tar ${NAME}:${TAG}

  echo "begin to scp to ${REMOTE_HOST}"
  scp "$1".tar kuan:/tmp

  echo "delete "$1".tar"
  rm "$1".tar

  # shellcheck disable=SC1009
  ssh kuan \
  "cd ../tmp && \
   docker stop web_user && \
   docker rm web_user && \
   docker rmi ${NAME}:${TAG} && \
   docker load -i "$1".tar && \
   docker run -dp NUMBER:NAME
}

build_service "$1"


