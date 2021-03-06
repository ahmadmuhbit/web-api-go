package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:8889)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	/* CRUD

	===========
	CREATE Data
	===========

	book := book.Book{}
	book.Title = "Halo Koding"
	book.Price = 120000
	book.Discount = 15
	book.Rating = 4
	book.Description = "Berkenalan dengan koding. Ubah dunia menjadi lebih baik"

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error creating book record")
		fmt.Println("==========================")
	}*/

	var books []book.Book

	err = db.Debug().Where("rating = ?", 5).Find(&books).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error finding book record")
		fmt.Println("==========================")
	}

	for _, b := range books {
		fmt.Println("Title :", b.Title)
		fmt.Println("book object %v", b)
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBooksHandler)

	router.Run()

}
