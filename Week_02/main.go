

//1. #暴力， sort， sorted_str 相等 ？ O(NlogN)
//2.  hash, map -- > 统计每个字符的拼频次
// https://leetcode-cn.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    words := make([]int, 26)

    for _, v:= range s {
        words[int(v) - 97]++
    }

    for _, v:= range t {
        words[int(v) - 97]--
    }

    for _, v:= range words {
        if v != 0{
            return false
        }
    }
    
    return true
}

// 1. 两层循环
// 2. 用hash表存储下标
// https://leetcode-cn.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
    a := make([]int, 2)
    numsSize := len(nums)
    for i := 0; i < numsSize - 1; i++{
        for j := i + 1; j < numsSize; j++ {
            if(nums[i] + nums[j] == target){
                a[0], a[1] = i, j
                return a
            }
        }
    }
    return []int{0}
}


//https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
//前序遍历的变种题目

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    res := []int{}
    helper(root, &res)
    return res
}

func helper(root *Node, res *[]int){
    if root == nil {
        return
    }

    *res = append(*res, root.Val)
    for _, v:= range root.Children {
        helper(v, res)
    }
}