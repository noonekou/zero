
.PHONY: dev

dev:
	go run rpc/add/add.go -f rpc/add/etc/add.yaml
	go run rpc/check/check.go -f rpc/check/etc/check.yaml
	go run api/bookstore.go -f api/etc/bookstore-api.yaml
