
.PHONY: dev check_api add_api start_add_rpc_dev start_check_rpc_dev start_bookstore_api_dev

dev: start_add_rpc_dev start_check_rpc_dev start_bookstore_api_dev

start_add_rpc_dev:
	cd rpc/add && go run add.go -f etc/add.dev.yaml
	cd ../../

start_check_rpc_dev:
	cd rpc/check && go run check.go -f etc/check.dev.yaml
	cd ../../

start_bookstore_api_dev:
	cd api && go run bookstore.go -f etc/bookstore-api.dev.yaml
	cd ../

check_api:
	curl -i "http://localhost:8888/check?book=go-zero"

add_api:
	curl -i "http://localhost:8888/add?book=go-zeroa&price=10"