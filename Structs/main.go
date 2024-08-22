package main

import "fmt"

type Node struct {
	val  int
	node *Node
}

func (N *Node) changeNodeVal(val int){
	N.val = val;
}

func main() {
	temp := Node{val: 5, node: nil};
	root := Node{val: 1, node: &temp}
	fmt.Println(root)
	root.changeNodeVal(2);
	root.node.changeNodeVal(3)
	fmt.Println(root.val,*root.node)
	fmt.Println(root)
}
