//387. 字符串中的第一个唯一字符
// 1. brute-force
// i 枚举所有字符
// 	j 枚举 i 后面的所有字符 // 找重复

// 返回 i 的索引

// O(N^2)

// 2. map（hashmap O(1), treemap O(logN)）
// O(N) or O(NlogN)

// 3. hash table 求字符的数量

func firstUniqChar(s string) int {
    for i := 0; i < len(s); i++ {
        for j := 0; j < len(s); j++ {
            if i != j && s[i] == s[j] {
                break
            }
            if j == len(s) - 1 { // 和最后一位位置长度一样，说明遍历到尾部都没有找到重复的元素
                return i
            }
        }
    }
    return -1
}

// 8. 字符串转换整数 (atoi)
//这道题 考验的是你的基本功， 练习小程序的能力非常重要
func myAtoi(str string) int {
    // 去掉首尾空格
    str = strings.TrimSpace(str)
    ans, sign := 0, 1

    for i, v := range str {
        if v >= '0' && v <= '9' {
            ans = ans * 10 + int(v - '0')
        } else if v == '-' && i == 0 {
            sign = -1
        } else if v == '+' && i == 0 {
            sign = 1
        } else{
            break
        }

    // 数值最大检测
    if ans > math.MaxInt32 {
        if sign == -1 {
            return math.MinInt32
        }
        return math.MaxInt32
    }
}
    return sign * ans
}

//91. 解码方法
// 0-9 dp[i-1]
// 10 - 26 dp[i-2]
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

//438. 找到字符串中所有字母异位词
// 滑动窗口做法， 套用模板
func findAnagrams(s string, p string) []int {
    window := make(map[byte]int)

     // 初始化 待匹配数据的数据状态
    need := make(map[byte]int)
    need_value := 0
    for v := range p { need[p[v]] ++  }
    for _ = range need { need_value ++}

    // 滑动窗口状态
    left, right := 0,0
    // 窗口与目标字符串对比状态
    valid := 0

    // 返回结果
    res := []int{}
    for right < len(s){
        c := s[right]
        right ++
        // 新增的字符在目标结果中
        if _,ok := need[c];ok{
            window[c]++
            if window[c] == need[c]{
                valid ++
            }
        }
        
        for right - left >= len(p) {
            // 找到目标更新结果
            if need_value == valid{
                res = append(res, left)
            }

            d := s[left]
            left ++

            // 如果删除的字符在目标字符串里面 则变更状态
            if _, ok := need[d]; ok{
                if window[d] == need[d]{
                    valid --
                }
                // 每次如果命中目标都应该减去window中的状态
                window[d]--
            }

        }

    }

    return res
}