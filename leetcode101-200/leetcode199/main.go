package main

 // Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 
// BFS
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := []int{}
	for len(queue) != 0 {
		l := len(queue)
		result = append(result, queue[l-1].Val)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return result
}

// DFS
var result []int
func rightSideView2(root *TreeNode) []int {
	result = make([]int, 0)
	dfs(root, 0)
	return result
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return 
	}

	if len(result) <= level {
		result = append(result, root.Val)
	}
	dfs(root.Right, level+1)
	dfs(root.Left, level+1)
}
