package main

import "fmt"

// Concrete Prototype: Folder
type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	// Create a new Folder
	clonedFolder := &Folder{name: f.name}
	var tempChildren []Inode
	// Clone each child node recursively
	for _, child := range f.children {
		tempChildren = append(tempChildren, child.clone())
	}
	clonedFolder.children = tempChildren
	return clonedFolder
}
