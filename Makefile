.SILENT:
.PHONY:

APP=cmd/app/main.go
PROJECT=learn/easyrest


app:
	swag init -g $(APP)
	go run $(APP)

fmt:
	go mod tidy
	go fmt $(PROJECT)/...
	# swag fmt .
