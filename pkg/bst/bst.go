package bst

type TreeNode struct {
	Val   float64
	Left  *TreeNode
	Right *TreeNode
}

var head TreeNode

func nilNode(node *TreeNode) bool {
	return node.Val == 0 && node.Left == nil && node.Right == nil
}

func addTreeNode(val float64, node *TreeNode) {
	newNode := TreeNode{Val: val, Left: nil, Right: nil}
	if node == nil {
		return
	}
	if val < node.Val {
		if node.Left != nil {
			addTreeNode(val, node.Left)
		} else {
			node.Left = &newNode
		}
	} else {
		if node.Right != nil {
			addTreeNode(val, node.Right)
		} else {
			node.Right = &newNode
		}
	}
}

func inorderTraversal(node *TreeNode, sortedOrderList *[]float64) {
	if node == nil {
		return
	}
	inorderTraversal(node.Left, sortedOrderList)

	*sortedOrderList = append(*sortedOrderList, node.Val)

	inorderTraversal(node.Right, sortedOrderList)
}

func SortedOrder() []float64 {
	var sortedOrder []float64
	inorderTraversal(&head, &sortedOrder)
	return sortedOrder
}

func AddNode(val float64) {
	if nilNode(&head) {
		head = TreeNode{Val: val, Left: nil, Right: nil}
	} else {
		addTreeNode(val, &head)
	}

}
