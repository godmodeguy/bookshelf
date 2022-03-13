.SILENT:
.PHONY:

APP=cmd/main.go
PROJECT=bookshelf


app:
	swag init -g $(APP)
	go run $(APP)

fmt:
	go mod tidy
	go fmt $(PROJECT)/...
	# swag fmt .
