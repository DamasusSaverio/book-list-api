package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	//Create
	// bookRequest := book.BookRequest{
	// 	Title: "Gundam",
	// 	Price: "200000",
	// }

	// bookService.Create(bookRequest)
	// FIND All
	// books, err := bookRepository.FindAll()
	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	//FIND by ID
	// book, err := bookRepository.FindByID(2)
	// fmt.Println("Title :", book.Title)

	//CREATE
	// book := book.Book{
	// 	Title:       "One Piece vol 79",
	// 	Description: "Rare book",
	// 	Price:       750000,
	// 	Rating:      4,
	// 	Discount:    0,
	// }

	// bookRepository.Create(book)

	//CRUD

	//CREATE data

	// book := book.Book{}
	// book.Title = "Atomic Habits"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 4
	// book.Description = "Ini adalah buku tentang self development"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")

	// }

	// READ data

	// var books []book.Book

	// err = db.Debug().Where("rating = ?", "5").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	//UPDATE data
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Man Tiger (Revised Edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("==========================")
	// }

	//DELETE data
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error deleting book record")
	// 	fmt.Println("==========================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}
