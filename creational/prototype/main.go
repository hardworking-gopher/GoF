package main

import "fmt"

func main() {
	// Create an initial prototype folder structure
	folder1 := &Folder{
		name: "Folder1",
		children: []Inode{
			&File{name: "File1-1"},
			&File{name: "File1-2"},
		},
	}

	fmt.Println("--- Original Structure ---")
	folder1.print("  ")

	// Clone the entire folder structure
	folder2 := folder1.clone()

	// You can now treat folder2 as a completely new object
	// For demonstration, let's cast it back to a Folder to modify it
	if cloned, ok := folder2.(*Folder); ok {
		cloned.name = "Folder2 (Cloned)"
		// Add a new file to the cloned folder to show it's independent
		cloned.children = append(cloned.children, &File{name: "File2-1 (New)"})
	}

	fmt.Println("\n--- Cloned and Modified Structure ---")
	folder2.print("  ")

	fmt.Println("\n--- Original Structure (Unchanged) ---")
	folder1.print("  ")
}
