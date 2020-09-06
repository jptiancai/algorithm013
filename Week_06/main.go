
//64. 最小路径和
// 老套路就是把遍历数组，找出上侧元素和左侧元素的最小值加上当前值，即可得到最小代价遍历图
func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    // 以下三个循环其实可以合并成在一起，但是为了代码的清晰和易于理解，还是分开写比较好
    for i := 1; i < m; i++ {
        grid[i][0] += grid[i - 1][0]
    }

    for j := 1; j < n; j++ {
        grid[0][j] += grid[0][j - 1]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])
        }
    }
    return grid[m - 1][n - 1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// 91. 解码方法
// 去除这些限制条件，此题就是爬楼梯的问题了，一次可以爬一步，也可以爬两步，问有多少中方式到达终点。
func numDecodings(s string) int {
    n := len(s)
    if n <= 0{
        return 0
    }
    dp := make([]int, n + 1)
    dp[0] = 1
    if s[0] == '0' {
        dp[1] = 0
    } else {
        dp[1] = 1
    }

    for i := 2; i <= n; i++ {
        first, _ := strconv.Atoi(s[i - 1 : i])
        second, _ := strconv.Atoi(s[i - 2 : i])
        if first >= 1 && first <= 9 {
            dp[i] += dp[i - 1]
        }
        if second >= 10 && second <= 26 {
            dp[i] += dp[i - 2]
        }
    }
    return dp[n]
}