**_PROBLEM FROM LEETCODE_**

> Given the `root` of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.
>
> As a reminder, a *binary search tree* is a tree that satisfies these constraints:
>
> - The left subtree of a node contains only nodes with keys **less than** the node's key.
>
> - The right subtree of a node contains only nodes with keys **greater than** the node's key.
>
> - Both the left and right subtrees must also be binary search trees.
>
>   ```
>   Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
>   Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
>   ```
>
>   **Example 2:**
>
>   ```
>   Input: root = [0,null,1]
>   Output: [1,null,1]
>   ```
>
>   **Example 3:**
>
>   ```
>   Input: root = [1,0,2]
>   Output: [3,3,2]
>   ```
>
>   **Example 4:**
>
>   ```
>   Input: root = [3,2,4,1]
>   Output: [7,9,4,10]
>   ```
>
>    
>
>   **Constraints:**
>
>   - The number of nodes in the tree is in the range `[0, 104]`.
>   - `-104 <= Node.val <= 104`
>   - All the values in the tree are **unique**.
>   - `root` is guaranteed to be a valid binary search tree.
>
>   **Note:** This question is the same as 1038: https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
>
>   **Example 1:**![img](https://assets.leetcode.com/uploads/2019/05/02/tree.png)



```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }    
    traverse(root, 0)    
    return root     
}

func traverse(node *TreeNode, rightSum int) int{
    var rightTree, leftTree, originalVal int
    originalVal = node.Val
    
    if node.Right != nil {
      // Get the sum of the nodes on the right
        rightTree = traverse(node.Right, rightSum)
    }
 
    if node.Left != nil {
      // Get the sum of the nodes on the left
        leftTree = traverse(node.Left, rightTree + originalVal + rightSum)
    }
    // Change current node value
    node.Val += rightTree + rightSum    
  	
  	// return value to be added
    return originalVal + rightTree + leftTree
}
```

