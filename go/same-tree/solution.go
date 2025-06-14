package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := [][2]*TreeNode{{p, q}}

	for len(stack) > 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		n1, n2 := pair[0], pair[1]

		if n1 == nil && n2 == nil {
			continue
		}

		if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		stack = append(stack, [2]*TreeNode{n1.Left, n2.Left})
		stack = append(stack, [2]*TreeNode{n1.Right, n2.Right})
	}

	return true
}
