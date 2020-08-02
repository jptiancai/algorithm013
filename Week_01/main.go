package main

// https://leetcode-cn.com/problems/two-sum/
// 两层循环
func twoSum(nums []int, target int) []int {
    if nums == nil {
        return []int{0}
    }
    a := make([]int, 2)
    numsSize := len(nums)
    for i := 0; i < numsSize - 1; i++ {
        for j := i + 1; j < numsSize; j++ {
            if nums[i] + nums[j] == target {
                a[0], a[1] = i, j
                return a
            }
        }
    }
    return []int{0}
}

// https://leetcode-cn.com/problems/move-zeroes/
// 时间复杂度O(n), 空间复杂度O(1)
// j++这个的位置写在了判断非零元素的外面，这样移动j指针就没有用处了
// 把核心逻辑写了三遍， 写完感觉心里很舒畅
// 看着自己的提交记录，5分钟前，1分钟前，几秒前
func moveZeroes(nums []int)  {
    // 边界检查
    if nums == nil {
		return
	}

    // 记录非零元素位置
	j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
}

