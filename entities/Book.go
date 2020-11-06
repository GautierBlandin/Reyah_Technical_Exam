/*
This file defines the Book entity.
 */
package entities

type Book struct {
	ID     int
	Title  string
	Author string
	Summary string
	ISBN   string
}