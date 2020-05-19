package Week_02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		l := len(stack)
		root = stack[l-1]
		stack = stack[:l-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
