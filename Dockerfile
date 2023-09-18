FROM golang:1.21 as builder

WORKDIR /src

COPY . .

ENV GOPROXY=https://goproxy.cn,direct

RUN rm -rf dev.go go.work go.work.sum ; go mod tidy
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /src/app


FROM alpine:latest as compress

WORKDIR /

COPY --from=builder /src/app /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories 
RUN  apk add --no-cache upx ca-certificates tzdata \
  && upx -5 app -o _upx_app \
  && mv -f _upx_app app

FROM busybox:stable-glibc

COPY --from=compress /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=compress /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=compress /app /app/storage
COPY ./config.prod.toml /app/storage.toml

ENTRYPOINT [ "/app/storage" ]