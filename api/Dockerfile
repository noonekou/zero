FROM golang:1.19-alpine
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/bookstore-api /app/bookstore.go

EXPOSE 8888
CMD ["/app/bookstore-api", "-f", "/app/etc/bookstore-api.yaml"]
