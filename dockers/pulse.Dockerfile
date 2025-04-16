FROM golang:1.22.1 as builder
WORKDIR /data
COPY ./ ./
RUN go env -w GO111MODULE="on" && go env -w GOPROXY="https://goproxy.cn" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /data/server ./app/cmd/main.go

FROM ubuntu:22.10
WORKDIR /data
COPY --from=builder /data/server /data/server
RUN ls -al
RUN mkdir -p /data/log
ENV INST_NO=0
ENV BEGIN_PORT=8000
ENV SERVICE=umdp
ENV LOCATION=CN
ENV LANG=C.UTF-8
ENTRYPOINT exec /data/server -c /data/config/config.yaml