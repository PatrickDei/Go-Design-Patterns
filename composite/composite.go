package composite

import "fmt"

func ShowcaseComposite() {
	fmt.Println("\nComposite pattern")

	file1 := File{
		name: "file1.go",
	}
	file2 := File{
		name: "file2.go",
	}

	folder := Folder{
		name:     "composite/",
		elements: []Element{&file1},
	}

	folder.add(&file2)

	file1.search("type File")
	file2.search("type Folder")
	folder.search("file1.go")
}
