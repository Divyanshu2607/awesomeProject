package matrix

import "fmt"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	PreOrder(root.Left)
	PreOrder(root.Right)
}
func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	fmt.Println(root.Value)
	PostOrder(root.Right)

}
