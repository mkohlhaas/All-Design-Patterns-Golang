package main

import "fmt"

//////////////// Interface ///////////////////////

type component interface {
	search(string)
}

//////////////// Compoonents /////////////////////

////////////////// Folder ////////////////////////

type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s.\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

// can add files and folders
func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

/////////////////// File /////////////////////////

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s.\n", keyword, f.name)
}

////////////// Main //////////////////////////////

func main() {
  // create folder hierarchy:
  // folder2
  //   folder1
  //      file1
  //   file2
  //   file3
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
	folder1 := &folder{
		name: "Folder1",
	}
	folder1.add(file1)
	folder2 := &folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

  // search in folder hierarchy recursively
	folder2.search("rose")
}
