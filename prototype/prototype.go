package prototype

import "fmt"

func ShowcasePrototype() {
	fmt.Println("\nPrototype pattern")
	f1 := File{name: "File1"}
	f2 := File{name: "File2"}
	f3 := File{name: "File3"}

	folder2 := Folder{
		children: []Node{&f1, &f2, &f3},
		name:     "Folder2",
	}

	cloneFolder := folder2.clone()

	cloneFolder.print("  ")

	folder2.print("  ")
}
