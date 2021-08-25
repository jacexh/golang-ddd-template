FROM golang:1.17 as builder
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
    && go build -ldflags "-X main.version=$REVISION" -o {{.BinFile}} cmd/main.go

FROM debian:buster
WORKDIR /app
VOLUME /app/conf
COPY --from=builder /go/src/{{.BinFile}} .
COPY --from=builder /go/src/conf/* ./conf/
RUN set -e \
    && sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list \
    && sed -i "s/security.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list \
    && apt update -yqq \
    && apt install -yqq ca-certificates \
    && apt clean autoclean \
    && apt autoremove -yqq \
    && rm -rf /var/lib/apt/lists/*
EXPOSE 8080
CMD ["/app/{{.BinFile}}"]