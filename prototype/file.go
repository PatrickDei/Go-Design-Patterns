package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation, f.name)
}

func (f *File) clone() Node {
	return &File{name: f.name}
}
