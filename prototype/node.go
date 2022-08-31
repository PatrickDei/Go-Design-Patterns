package prototype

type Node interface {
	print(string)
	clone() Node
}
