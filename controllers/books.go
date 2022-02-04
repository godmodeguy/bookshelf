package controllers

import (
	"learn/easyrest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// DTO section
// Data Transferr Objects
// Only for transfet data between layers

type CreateBookInput struct {
	Title 	string `json:"title" binding:"required"`
	Author 	string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title 	string `json:"title"`
	Author 	string `json:"author"`
}

// @Summary      Get books
// @Description  return all books in json
// @Tags         books
// @Produce      json
// @Router       /books  [get]
func FindBooks(c *gin.Context)  {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}


// @Summary      Create book
// @Description  write book info into db
// @Tags         books
// @Param        input   body  CreateBookInput  true  "book title"
// @Accept       json
// @Produce      json
// @Router       /books  [post]
func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}


// @Summary      Find book
// @Description  search for book with given id
// @Tags         books
// @Param        id      path   int     true   "Book ID"
// @Produce      json
// @Router       /books/{id}  [get]
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nothing found"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"data": book})
}


// @Summary      Update book
// @Description  change info about book
// @Tags         books
// @Param        id  path  int  true  "Book ID"
// @Param        input   body  UpdateBookInput  false  "book title"
// @Accept       json
// @Produce      json
// @Router       /books/{id}  [put]
func  UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}


// @Summary      Delete book
// @Description  remove book from db with given id
// @Tags         books
// @Param        id  path  int  true  "Book ID"
// @Produce      json
// @Router       /books/{id}  [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}