package main

import (
	api "bookshelf/pkg/api/v1"
)

func main() {
	handler := api.NewHandler()
	r := handler.InitRoutes()
	r.Run(":1337")
}
