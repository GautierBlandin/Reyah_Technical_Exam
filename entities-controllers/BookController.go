/*
	This files connects our program to the database. It allows database initialization (hostname,
	password and database address must be provided), and CRUD operations on Books stored in the
	database
 */

package entities_controllers

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"technical-test-reyah/entities"
)

const hostname = "admin"
const password = "admin"
const db_address = "127.0.0.1:3306"
const db_name = "reyah_technical_exam_db"

//Initialize the database with db_name and table books
func InitializeDatabase(){
	//db is an object allowing us to execute queries in our database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", hostname, password, db_address))
	if err != nil{
		panic(err)
	}

	//Create database the database if it does not exist yet
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + db_name)

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", hostname, password, db_address, db_name))
	if err != nil{
		panic(err)
	}

	//Create table if it does not exist yet
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`.`books` ( `ID` INT NOT NULL AUTO_INCREMENT , `Title` TEXT NOT NULL , `Author` TEXT NULL DEFAULT NULL, `Summary` TEXT NULL DEFAULT NULL , `ISBN` TEXT NULL DEFAULT NULL , PRIMARY KEY (`ID`))", db_name ))
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	db.Close()
}

//Execute the given query and return the result
func queryDatabase(query string) *sql.Rows{

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", hostname, password, db_address, db_name))
	if err != nil{
		panic(err)
	}

	// defer the close until after the function has finished
	defer db.Close()

	queryResult, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	return queryResult
}

//Fetch book with specified ID from the database
func GetBook(ID int) (entities.Book, error){
	queryResult := queryDatabase(fmt.Sprintf("SELECT * FROM BOOKS WHERE ID = '%d'", ID))
	defer queryResult.Close()

	var book = entities.Book{}

	queryResult.Next()
	err := queryResult.Scan(&book.ID, &book.Title, &book.Author, &book.Summary, &book.ISBN)
	if err != nil {
		return entities.Book{}, errors.New("No book found for specified ID")
	}

	return book, nil
}

//Fetch all books from the database
func GetBooks() ([]entities.Book, error){
	queryResult := queryDatabase("SELECT * FROM BOOKS")
	defer queryResult.Close()

	//Declaring returned array
	var books []entities.Book

	//Append each book to the array of books
	for queryResult.Next(){
		var book = entities.Book{}
		err := queryResult.Scan(&book.ID, &book.Title, &book.Author, &book.Summary, &book.ISBN)

		//Return partial list in case of error, and returns the error
		if(err != nil){
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

//Add a book to the database
func CreateBook(book entities.Book) {
	queryResult := queryDatabase(fmt.Sprintf("INSERT INTO BOOKS (Title, Author, Summary, ISBN) VALUES ('%s', '%s', '%s', '%s')", book.Title, book.Author, book.Summary, book.ISBN))
	queryResult.Close()
}

//Delete book with specified ID from the database
func DeleteBook(ID int){
	queryResult := queryDatabase(fmt.Sprintf("DELETE FROM BOOKS WHERE ID = %d", ID))
	queryResult.Close()
}

//Update given book in the database (ID field MUST NOT be 0)
func UpdateBook(book entities.Book) error{

	if book.ID == 0{
		return errors.New("Book has no ID")
	}

	queryResult := queryDatabase(fmt.Sprintf("UPDATE BOOKS SET Title = '%s', Author = '%s', Summary = '%s', ISBN = '%s' WHERE ID = '%d' ", book.Title, book.Author, book.Summary, book.ISBN, book.ID))
	queryResult.Close()
	return nil

}