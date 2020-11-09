The API folder contains all the code that I have written for this project.
The github.com folder contains libraries used in the project.
I used a MySQL database, but any SQL database can be used, it is simply necessary to change the driver in /API/entities-controllers/BookController.go

The API is used by running /API/api.go. It listens on port 8080, and can be tested by sending HTTP Requests with softwares such as Postman. 
Before use, hostname, password and database address must be updated in /API/entities-controllers/BookController.go, or the API will not be able to connect to your database. The database itself and the books table do not need to be created beforehand, they will be created automatically at first program launch.

The available functionalities are :
	GET /book-management/books/:id : Returns the specified book in the response body
	GET /book-management/books : Returns all the books in the response body
	POST /book-management/books : Creates a new book. Keys : title, author, summary, isbn
	DELETE /book-management/books/:id : Deletes book with specified ID
	PUT /book-management/books/:id : Updates the specified book. Keys : title, author, summary, isbn

Ideas for further expansion : creating a dedicated go file to connect to and manage the database (currently done directly inside BookController.go), using a .env file for database constants.