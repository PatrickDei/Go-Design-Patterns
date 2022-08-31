package composite

import "fmt"

type Folder struct {
	name     string
	elements []Element
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching for %v in %v\n", keyword, f.name)
}

func (f *Folder) getName() string {
	return f.name
}

func (f *Folder) add(e Element) {
	f.elements = append(f.elements, e)
}
