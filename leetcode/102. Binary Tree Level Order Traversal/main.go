/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    result := [][] int {}
    if root == nil {
        return result
    }
    queue := []*TreeNode{root}
    

    for len(queue)>0{
        level := []int{}
        qLen := len(queue)  // lenght to iterate various levels
        for i := 0; i<qLen;i++{
            node := queue[0]        // Remove node from queue
            queue = queue[1:]
            level = append(level, node.Val)     // capture the value from each node
            if node.Left != nil{
                queue = append(queue,node.Left)     // If Left add to queue
            }
            if node.Right != nil{
                queue = append(queue,node.Right)    // If Right add to queue
            }
        }

        result = append(result, level)
    }
    return result
}

// Input:
// root = [3,9,20,null,null,15,7]

// Output:
// [[3],[9,20],[15,7]]


