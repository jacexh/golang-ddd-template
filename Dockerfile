FROM golang:1.14 as builder
WORKDIR /go/src
COPY . /go/src
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
RUN set -e \
    && sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list \
    && sed -i "s/security.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list \
    && apt update -y \
    && apt install -y git \
    && REVISION=`git rev-list -1 HEAD` \
    && go build -ldflags "-X main.version=$REVISION" -o {{.BinFile}} -tags=jsoniter cmd/main.go

FROM debian:buster
WORKDIR /app
COPY --from=builder /go/src/app .
COPY --from=builder /go/src/cmd/config.yml .
EXPOSE 8080
CMD ["/app/{{.BinFile}}"]