package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	Val int
	Left *Node
	Right *Node
}

func (tree *Tree) insertNode(val int)  {
	newNode:= &Node{Val: val}
	if tree.root == nil{
		tree.root= newNode
		return
	}
	current:= tree.root
	for {
		if val < current.Val{
			if current.Left == nil{
				current.Left= newNode
				return
			}
			current= current.Left
		}else{
			if current.Right == nil{
				current.Right= newNode
				return
			}
			current= current.Right

		}
	}
}

func (tree *Tree) find(value int) bool {
	current:= tree.root

	for current != nil{
		if value < current.Val{
			current= current.Left
		}else if value > current.Val {
			current= current.Right
		}else {
			return true
		}
	}
	return false
}
func (tree *Tree) PreOrder(root *Node)  {
	if root == nil{
		return
	}
	fmt.Println(root.Val)
	tree.PreOrder(root.Left)
	tree.PreOrder(root.Right)

}

func (tree *Tree) inOrder(root *Node)  {
	if root == nil{
		return
	}
	tree.inOrder(root.Left)
	fmt.Println(root.Val)
	tree.inOrder(root.Right)
}

func (tree *Tree) sum(root *Node)int  {
	sum:= 0
	if root == nil{
		return 0
	}
	for {
		if root.Left != nil{
			sum += root.Val
		}
		if root.Right != nil{
			sum += root.Val
		}
		break
	}

	//tree.PreOrder(root.Left)

	//tree.PreOrder(root.Right)
	return sum
}

func main() {
	tree:= &Tree{}

	tree.insertNode(27)
	tree.insertNode(14)
	tree.insertNode(35)
	tree.insertNode(42)
	tree.insertNode(10)
	tree.insertNode(19)
	tree.insertNode(31)
	tree.sum(tree.root)
	//tree.inOrder(tree.root)
	//fmt.Println("---------------------")
	//tree.PreOrder(tree.root)
	//fmt.Println(tree.find(37))

}
