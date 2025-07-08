package main

// Prototype: Interface for cloning
type Inode interface {
	print(indentation string)
	clone() Inode
}
