package utils

import (
	"fmt"
	"log"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Name     string
	Kind     string
	Value    int
	Parent   *Node
	Children map[string]*Node
}

func (t *Tree) New() *Tree {
	if t.Root == nil {
		t.Root = &Node{Value: 0, Parent: nil, Name: "root", Kind: "dir", Children: make(map[string]*Node)}
	} else {
		log.Fatal("there already is a root node")
	}

	return t
}

func (t *Tree) print(tree Node, indent string) {
	if tree.Kind == "dir" {
		fmt.Println(indent+tree.Name, "(", tree.Value, ")")
	} else {
		fmt.Println(indent + tree.Name)
	}

	indent = indent + " "
	for _, val := range tree.Children {
		t.print(*val, indent)
	}
}

func (t *Tree) Print() {
	t.print(*t.Root, "")
}

func (n *Node) Insert(node *Node) {
	// If not a dir no insert
	if n.Kind != "dir" {
		log.Fatalln("cant insert a file into a file")
	}

	// This is super un optimized but we add the value of the new node to each parent of the current node
	if node.Value > 0 {
		temp := n

		for temp != nil {
			temp.Value += node.Value
			temp = temp.Parent
		}
	}

	n.Children[node.Name] = node
}

func (n *Node) Delete(node *Node) {
	n.Children[node.Name] = nil
}
