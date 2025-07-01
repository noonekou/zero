FROM golang:1.19-alpine
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/check /app/check.go

EXPOSE 8081
CMD ["/app/check", "-f", "/app/etc/check.yaml"]
