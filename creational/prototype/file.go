package main

import "fmt"

// Concrete Prototype: File
type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	// Create a new File instance with the same name
	return &File{name: f.name}
}
