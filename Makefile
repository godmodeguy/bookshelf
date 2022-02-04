
run:
	swag init -g cmd/main.go
	go run cmd/main.go

fmt:
	go mod tidy
	go fmt learn/easyrest/...
	# swag fmt .

