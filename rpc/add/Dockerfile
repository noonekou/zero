FROM golang:1.19-alpine
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/add /app/add.go

EXPOSE 8080
CMD ["/app/add", "-f", "/app/etc/add.yaml"]
