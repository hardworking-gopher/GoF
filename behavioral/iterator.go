package main

import "fmt"

// --- Product (just for context) ---
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("'%s' by %s", b.Title, b.Author)
}

// --- 1. Iterator (Interface) ---
type BookIterator interface {
	HasNext() bool
	Next() *Book // Returns the next book
	Reset()      // Resets the iterator to the beginning
}

// --- 4. Concrete Aggregate ---
// A custom collection of books.
type BookCollection struct {
	books []*Book // Our internal representation (a slice)
}

func NewBookCollection() *BookCollection {
	return &BookCollection{
		books: make([]*Book, 0),
	}
}

func (bc *BookCollection) AddBook(book *Book) {
	bc.books = append(bc.books, book)
}

// --- 3. Aggregate (Interface) -- Implemented by BookCollection
// CreateIterator is the "factory method" for creating an iterator.
func (bc *BookCollection) CreateIterator() BookIterator {
	return &BookCollectionIterator{
		collection: bc,
		index:      0,
	}
}

// --- 2. Concrete Iterator ---
// Implements the BookIterator interface for BookCollection.
type BookCollectionIterator struct {
	collection *BookCollection // Reference to the aggregate
	index      int             // Current position
}

func (bci *BookCollectionIterator) HasNext() bool {
	return bci.index < len(bci.collection.books)
}

func (bci *BookCollectionIterator) Next() *Book {
	if bci.HasNext() {
		book := bci.collection.books[bci.index]
		bci.index++ // Move to the next element
		return book
	}
	return nil // Or return an error, depending on desired behavior
}

func (bci *BookCollectionIterator) Reset() {
	bci.index = 0
}

// --- Client Code ---
func main() {
	// Create a concrete aggregate
	library := NewBookCollection()
	library.AddBook(&Book{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien"})
	library.AddBook(&Book{Title: "Pride and Prejudice", Author: "Jane Austen"})
	library.AddBook(&Book{Title: "1984", Author: "George Orwell"})
	library.AddBook(&Book{Title: "To Kill a Mockingbird", Author: "Harper Lee"})

	// Get an iterator from the aggregate
	// The client only interacts with the BookIterator interface.
	// It doesn't know that internally BookCollection uses a slice.
	it := library.CreateIterator()

	fmt.Println("--- Iterating through the library (forward) ---")
	for it.HasNext() {
		book := it.Next()
		fmt.Printf("Reading: %s\n", book)
	}

	fmt.Println("\n--- Resetting and iterating again ---")
	it.Reset() // Reset the iterator to the beginning
	for it.HasNext() {
		book := it.Next()
		fmt.Printf("Re-reading: %s\n", book)
	}

	// Example of another independent iteration
	fmt.Println("\n--- Another independent iteration ---")
	anotherIt := library.CreateIterator()
	anotherIt.Next()                                                           // Advance this iterator once
	fmt.Printf("First book from independent iterator: %s\n", anotherIt.Next()) // Get second book

	fmt.Printf("Original iterator (it) is still at the end: HasNext() = %t\n", it.HasNext())
}
