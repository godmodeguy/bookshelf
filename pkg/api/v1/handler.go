package api_v1

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		collections := api.Group("/lists")
		{
			collections.POST("/", h.createCollection)
			collections.GET("/", h.getCollection)
			collections.GET("/:id", h.getCollection)
			collections.PUT("/:id", h.updateCollection)
			collections.DELETE("/:id", h.deleteCollection)
		}

		books := api.Group("/:id/item")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getBook)
			books.GET("/:id", h.getBook)
			books.PUT("/:id", h.updateBook)
			books.DELETE("/:id", h.deleteBook)
		}
	}

	return r
}
