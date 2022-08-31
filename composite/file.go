package composite

import "fmt"

type Element interface {
	search(keyword string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for %v in %v\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
