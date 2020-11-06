
/*
	This files defines the API available at port 8080 on localhost

	Available functionalities :
	GET /book-management/books/:id : returns the specified book in the response body
	GET /book-management/books : returns all the books in the response body
	POST /book-management/books : Create a new book. Keys : title, author, summary, isbn
	DELETE /book-management/books/:id : delete book with specified ID
	PUT /book-management/books/:id : updates the specified book. Keys : title, author, summary, isbn
*/


package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technical-test-reyah/entities"
	"technical-test-reyah/entities-controllers"
)

func main() {

	Router := gin.Default()
	entities_controllers.InitializeDatabase()

	//Returns the book with specified ID in the response body
	Router.GET("/book-management/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			c.String(http.StatusBadRequest, "Syntax error in request")
			return
		}

		book, err := entities_controllers.GetBook(id)
		if err != nil {
			c.String(404, "No book for given ID")
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("ID : %d, Title : %s, Author : %s, Summary : %s, ISBN : %s",book.ID, book.Title, book.Author, book.Summary, book.ISBN))

	})


	//Returns all the books in the response body
	Router.GET("/book-management/books", func(c *gin.Context) {

		var responseBody string
		books, err := entities_controllers.GetBooks()

		if err != nil {
			c.String(404, "Error when fetching books")
			return
		}

		for _, book := range(books){
			responseBody += fmt.Sprintf("ID : %d, Title : %s, Author : %s, Summary : %s, ISBN : %s \n",book.ID, book.Title, book.Author, book.Summary, book.ISBN)
		}

		c.String(http.StatusOK, responseBody)

	})

	//Create a new book. Keys : title, author, summary, isbn
	Router.POST("/book-management/books", func(c *gin.Context) {
		//Query keys : title, summary, author, isbn

		//Declaring the book variable to be stored later
		var book entities.Book

		//Setting title and rejecting the request if the title is empty
		book.Title = c.Query("title")
		if book.Title == ""{
			c.String(http.StatusBadRequest, "Book must have a title")
			return
		}

		//Setting the other book fields
		book.Summary = c.Query("summary")
		book.Author = c.Query("author")
		book.ISBN = c.Query("isbn")

		//If the query is correct, store the book in the database

		entities_controllers.CreateBook(book)
		c.String(http.StatusCreated, "Created book")

	})

	//Delete the specified book
	Router.DELETE("/book-management/books/:id", func(c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			c.String(http.StatusBadRequest, "Syntax error in request")
			return
		}

		entities_controllers.DeleteBook(id)

		c.String(http.StatusOK, "Delete book with id %d", id)
	})

	//Update the specified book. Keys that do not exist will not be updated. Keys with empty value will be updated and field will become empty
	Router.PUT("/book-management/books/:id", func(c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			c.String(http.StatusBadRequest, "Syntax error in request")
			return
		}

		book, err := entities_controllers.GetBook(id)
		if err != nil{
			c.String(http.StatusNotFound, "Book doesn't exist")
			return
		}

		title, titleExists := c.GetQuery("title")
		if(titleExists){
			book.Title = title
		}

		author, authorExists := c.GetQuery("author")
		if authorExists{
			book.Author = author
		}

		summary, summaryExists := c.GetQuery("summary")
		if summaryExists{
			book.Summary = summary
		}

		isbn, isbnExists := c.GetQuery("isbn")
		if isbnExists{
			book.ISBN = isbn
		}

		err = entities_controllers.UpdateBook(book)

		c.String(http.StatusAccepted, "Book updated")
	})

	Router.Run(":8080")
}