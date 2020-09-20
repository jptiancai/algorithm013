//191. 位1的个数
// 循环次数是1的个数 的做法

// 把n最低位的1 变为0 : x & (x - 1)

// while(x > 0) {count ++; x = x & x(x - 1)}


// 循环32次的做法

// 1. for loop: 0 -- > 32, 判断1的个数
// 2. %2 + /2 ，打掉最后1位
// 3. &1， x = x >> 1; 
func hammingWeight(num uint32) int {
    sum := 0
    for num != 0  {
        sum ++
        num &= (num - 1)
    }
    return sum
}

// 231. 2的幂
// n & n - 1, 看它n的最低位1 打掉后 是否为0
func isPowerOfTwo(n int) bool {
    return n != 0 && (n & (n - 1)) == 0
}

// 190. 颠倒二进制位
// for loop

// i 和 32 - i位置的 二进制互换位置

func reverseBits(num uint32) uint32 {
    var ans uint32
    for i := 0; i < 32; i ++ {
        ans = (ans << 1) + (num & 1)
        num >>= 1
    }
    return ans >> 0
}

// 52. N皇后 II
// 位运算

var count int
func totalNQueens(n int) int {
    count = 0
    dfs(n, 0, 0, 0, 0)
    return count
}

func dfs(n, row, col, pie, na int) {
    if row >= n {
        count++
        return
    }

    bits := (^(col | pie | na)) & ((1 << uint(n)) - 1) // 获取所有可填空位
    for bits != 0 {
        bit := bits & -bits // 取位置排在最后的一个空位
        dfs(n, row + 1, col | bit, (pie | bit) << 1, (na | bit) >> 1) // 递归遍历下一行，列、撇、捺相应的增加
        bits = bits & (bits - 1) // 打掉最后位置上的1， 因为该位置已被占用
    }
}