package terminal_state

import (
	"fmt"
	"strings"
)

type Node struct {
	Parent   *Node
	Name     string
	Children map[string]*Node
	Size     int
}

func (n *Node) AsMap() map[string]*Node {
	m := map[string]*Node{}
	n.addAllToMapRec(m, "")
	return m
}

func (n *Node) addAllToMapRec(m map[string]*Node, prefix string) {
	full_path := fmt.Sprintf("%s/%s", prefix, n.Name)
	m[full_path] = n
	for _, c := range n.Children {
		c.addAllToMapRec(m, full_path)
	}
}

func (n *Node) AddSize(size int) {
	n.Size += size
	if n.Parent != nil {
		n.Parent.AddSize(size)
	}
}

func (n *Node) PrettyPrint() {
	n.prettyPrintRec(0)
}

func (n *Node) prettyPrintRec(indent int) {
	fmt.Printf("%s%s:%d\n", strings.Repeat(" ", indent), n.Name, n.Size)
	for _, c := range n.Children {
		c.prettyPrintRec(indent + 4)
	}
}
