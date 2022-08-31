package prototype

import "fmt"

type Folder struct {
	children []Node
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation, f.name)
	for _, c := range f.children {
		c.print(indentation + indentation)
	}
}

func (f *Folder) clone() Node {
	var tempChildren []Node
	for _, c := range f.children {
		childClone := c.clone()
		tempChildren = append(tempChildren, childClone)
	}

	return &Folder{
		children: tempChildren,
		name:     f.name,
	}
}
