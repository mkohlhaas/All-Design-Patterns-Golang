package main

import "fmt"

////////////////// Interface /////////////////////

type inode interface {
	print(string)
	clone() inode
}

////////////////// Prototypes ////////////////////

/////////////// File Prototype ///////////////////

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

/////////////// Folder Prototype /////////////////

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	clonedFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	clonedFolder.children = tempChildren
	return clonedFolder
}

///////////////// Main ///////////////////////////

func main() {
	// creating a "file system"
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
	subFolder := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}
	folder := &folder{
		children: []inode{subFolder, file2, file3},
		name:     "Folder",
	}

	// print "file system"
	fmt.Println("Printing folder hierarchy:")
	folder.print("  ")

	// clone "file system"
	clonedFolder := folder.clone()

	// print cloned "file system"
	fmt.Println("\nPrinting cloned folder hierarchy:")
	clonedFolder.print("  ")
}
