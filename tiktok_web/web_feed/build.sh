#!/bin/bash

# 该脚本可用于自动打包，上传到服务器，并重启指定服务
# 执行 chmod +x build.sh
# ./build.sh web_user

set -e

REMOTE_HOST="124.223.70.243"
TAG=1.0.0.0
NUMBER=6002
NAME="localhost:5000/web_feed"

build_service() {
  echo "name: ${1}"

  cd /Users/kuan525/Desktop/go-tiktok/tiktok_web/${1}

  echo "go mod vendor.........."
  go mod vendor

  echo "begin to build ${NAME}:${TAG}..........."
  docker build --platform linux/amd64 -t ${NAME}:${TAG} .

  echo "begin to save ${1}.tar"
  docker save -o ${1}.tar ${NAME}:${TAG}

  echo "begin to scp to ${REMOTE_HOST}"
  scp ${1}.tar kuan:/tmp

  echo "delete ${1}.tar"
  rm ${1}.tar

  # shellcheck disable=SC1009
  ssh kuan \
  "cd /tmp && \
   docker stop ${1} && \
   docker rm ${1} && \
   docker rmi ${NAME}:${TAG} && \
   docker load -i ${1}.tar && \
   docker run -itd -p ${NUMBER}:${NUMBER} --name ${1} ${NAME}:${TAG}"
#   docker run -itd --platform linux/arm64/v8 -dp 6001:6001 --name web_user localhost:5000/web_user:1.0.0.0
}

build_service ${1}


