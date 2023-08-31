package gee

import(
	"fmt"
)


type Node struct{
	pattern string
	part string
	children []*Node
	isWild bool
}

func (n *Node) matchChild(part string) *Node{

	for _, child := range n.children{
		if child.part == part || child.isWild{
			return child
		}
	}
	return nil
}

func (n *Node) matchChildren(part string) []*Node{
	children := make([]*Node, 0)
	for _, child := range n.children{
		if child.part == part || child.isWild{
			children = append(children, child)
		}
	}
	return children
}

func (n *Node) insert(pattern string, parts []string, height int){
	if len(parts) == height{
		n.pattern = pattern
		return
	}

	child := n.matchChild(parts[height])
	if child == nil{
		newChild := &Node{
			part: parts[height],
			isWild: parts[height][0] == '*' || parts[height][0] == ':',
		}
		n.children = append(n.children, newChild)
		child = newChild
	}

	child.insert(pattern, parts, height + 1)
}

func (n *Node) search(parts []string, height int) *Node{
	if len(parts) == height{
		fmt.Printf("%v\n", n.pattern)
		return n
	}

	children := n.matchChildren(parts[height])

	for _, child := range children{
		//fmt.Printf("%v\n", child.part)
		ret := child.search(parts, height + 1)
		if ret != nil{
			return ret
		}	
	}
	return nil
}