package books

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/vikalpacn/Library/authors"
	"github.com/vikalpacn/Library/publishers"
)

const fileName = "books.json"

type Book struct {
	ID        string
	Title     string
	Author    authors.Author
	Genre     string
	Publisher publishers.Publisher
	Language  string
}

type Books struct {
	Books []Book
}

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllBooks() Books {
	err := CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}

	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", fileName)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var books Books

	json.Unmarshal(byteValue, &books)

	return books
}

func (b Books) PrintAll() {
	for i := 0; i < len(b.Books); i++ {
		fmt.Println(b.Books[i].ToString())
	}
}

func (b Book) ToString() string {
	return "ID: " + b.ID + ", " + "Title: " + b.Title + ", " + "Author: " + b.Author.Name + ", " + "Genre: " + b.Genre + ", " + "Publisher: " + b.Publisher.Name + ", " + "Language: " + b.Language
}

func (b Books) GetBookByID(id string) Book {
	for i := 0; i < len(b.Books); i++ {
		if b.Books[i].ID == id {
			fmt.Printf("Book found: ID: %s\n Title: %s\n Author: %s\n Genre: %s\n Publisher: %s\n", b.Books[i].ID, b.Books[i].Title, b.Books[i].Author.Name, b.Books[i].Genre, b.Books[i].Publisher)
			return b.Books[i]
		}
	}

	return Book{}
}


func AddBook(book Book) {
	books := GetAllBooks()
	books.Books = append(books.Books, book)

	updatedBooksJSON, err := json.Marshal(books)
	if err != nil {
		log.Println("Error marshaling books:", err)
		return
	}

	err = ioutil.WriteFile(fileName, updatedBooksJSON, 0644)
	if err != nil {
		log.Println("Error writing books file:", err)
		return
	}

	fmt.Println("Book added successfully.")
}

func createBook(scanner *bufio.Scanner) Book {
	fmt.Println("Adding a new book:")
	fmt.Print("Enter the ID: ")
	scanner.Scan()
	id := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the author name: ")
	scanner.Scan()
	authorName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the genre: ")
	scanner.Scan()
	genre := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the publisher: ")
	scanner.Scan()
	publisher := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the language: ")
	scanner.Scan()
	language := strings.TrimSpace(scanner.Text())

	// Create and return the book object
	book := Book{
		ID:        id,
		Title:     title,
		Author:    authors.Author{Name: authorName},
		Genre:     genre,
		Publisher: publishers.Publisher{Name: publisher},
		Language:  language,
	}

	return book
}

func addBook(scanner *bufio.Scanner) {
	book := createBook(scanner)
	fmt.Println("Book added successfully!",book)
}


func RemoveBookByName(name string) {
	books := GetAllBooks()

	// Find the index of the book with the given name
	index := -1
	for i, book := range books.Books {
		if book.Title == name {
			index = i
			break
		}
	}

	// If the book was found, remove it from the list
	if index != -1 {
		books.Books = append(books.Books[:index], books.Books[index+1:]...)
	} else {
		fmt.Println("Book not found.")
		return
	}

	updatedBooksJSON, err := json.Marshal(books)
	if err != nil {
		log.Println("Error marshaling books:", err)
		return
	}

	err = ioutil.WriteFile(fileName, updatedBooksJSON, 0644)
	if err != nil {
		log.Println("Error writing books file:", err)
		return
	}

	fmt.Println("Book removed successfully.")
}
