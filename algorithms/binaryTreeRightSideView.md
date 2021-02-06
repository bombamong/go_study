### Binary Tree Right Side View

Given a binary tree, imagine yourself standing on the _right_ side of it, return the values of the nodes you can see ordered from top to bottom.

**Example:**

```
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
```

**MySolution**

> **Strategy**
>
> traverse both left and right sides of the tree until there are no more leaves.
>
> Compare length of left and right branches.
>
> If the right branch is longer than the left branch that means that there is nothing to worry about and should return the right branch
>
> However, if the left branch is longer, it means that there are nodes that are at deeper levels than the deepest rightmost node at the right branch which need to be into account.
> Therefore, any node that in the left branch that share the same depth as an existent node at the right branch should be overrriden by the rightside branch.
>
> Ex)
>
> ```  
>         1
>       2   3
>     5   6   4
>       7   8
>      9
> ```
>
> After both branches are combined together, the new branch should be returned.

```go
func rightSideView(root *TreeNode) []int {
    var answer []int
    if root == nil {
        return answer
    }
    var rightNodes, leftNodes []int
    leftNodes = rightSideView(root.Left)
    rightNodes = rightSideView(root.Right)
    answer = append(answer, root.Val)
    if len(rightNodes) > len(leftNodes) {
        answer = append(answer, rightNodes...)
    } else {
        for i := 0; i < len(rightNodes); i++ {
            leftNodes[i] = rightNodes[i]
        }
        answer = append(answer, leftNodes...)
    }
    return answer
}
```
