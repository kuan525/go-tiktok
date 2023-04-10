FROM golang:1.19 as builder

# Local
# 先执行go mod vendor，否则找不到外层的依赖
COPY . /Users/kuan525/Desktop/go-tiktok/tiktok_web/web_user

# GOSUMDB:关闭和sum.golang.org的交互  CGO_ENABLED:禁用CGO
# -trimpath 可以削减二进制文件的元数据路径  -ldflags "-w -s" 用于去除调试信息并缩小二进制文件的大小
RUN cd /Users/kuan525/Desktop/go-tiktok/tiktok_web/web_user && GOPROXY=https://goproxy.cn,direct GOSUMDB=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s" -o web_user cmd/main.go

FROM alpine

# 指定作者
LABEL maintainer="kuan525 <kuan_525@163.con>"
# 更改默认用户并以非特权用户身份运行应用程序
USER root

# 上海时区 语言美国 utf-8 本地时间设置为上海时区的系统时间 更改镜像网站 安装tzdata软件包，为容器中其他应用程序设置正确的时区提供支持
ENV TZ "Asia/Shanghai"
ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US.UTF-8
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo 'Asia/Shanghai' >/etc/timezone \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add tzdata

WORKDIR /workspace

# 指定源Docker映像
COPY --from=builder /go/go-tiktok/go-tiktok /workspace

ENTRYPOINT ["./web_user"]

