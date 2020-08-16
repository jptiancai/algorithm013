
// 236. 二叉树的最近公共祖先
//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 分别递归左子树、右子树
// 关于边界值的判断

// 左右子树都不为空， 说明为根节点
// 左结点为空， 说明最近的公共祖先在右节点

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

//105. 从前序与中序遍历序列构造二叉树
//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 关键找到根节点，然后分为左子树 和 右子树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    // 中序遍历找到根节点
    var root int
    for k, v := range inorder {
        if v == preorder[0] {
            root = k
            break
        }
    }
    // 放到数的左右位置
    return &TreeNode{
        Val:  preorder[0],
        //左右子树递归
        Left: buildTree(preorder[1 : root + 1], inorder[:root]),
        Right: buildTree(preorder[root + 1:], inorder[root + 1:]),
    }
    
}


// 77. 组合
//https://leetcode-cn.com/problems/combinations/
// 处理当前状态 和 下探到下一层 可以合并为一行， 代码更加的简洁

func combine(n int, k int) [][]int {
    result := make([][]int, 0)
    helper(1, n, k, []int{}, &result)
    return result
}

func helper(porinter, n, k int, current []int, result *[][]int) {
    if len(current) == k {
        *result = append(*result, append([]int{}, current...))
        return
    }

    for i := porinter; i <= n; i++ {
        helper(i + 1, n, k, append(current, i), result)
    }
}

// 46. 全排列
// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
    res := [][]int{}
    helper(nums, []int{}, &res)
    return res
}

func helper(nums []int, prev []int, res *[][]int) {
    if len(nums) == 0 {
        *res = append(*res, append([]int{}, prev...))
        return
    }

    for i := 0; i < len(nums); i ++ {
        helper(append(append([]int{}, nums[:i]...), nums[i + 1:]...),
                append(prev, nums[i]),
                res)
    }
}