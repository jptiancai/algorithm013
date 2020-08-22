//柠檬水找零
// https://leetcode-cn.com/problems/lemonade-change/
// 1.贪心算法

// 如果给5元，不用找
// 如果给10元，找5元，否则找不开
// 如果给20元，优先找10 + 5元， 再找5+5+5， 否则找不开

func lemonadeChange(bills []int) bool {
    five, ten := 0, 0

    for _, v := range bills {
        if v == 5 {
            five ++
        } else if v == 10 {
            if five >= 1 {
                five --
                ten ++
            } else {
                return false
            }
        } else if v == 20 {
            if five >= 1 && ten >= 1 {
                five --
                ten --
            } else if five >= 3 {
                five -= 3
            } else {
                return false
            }
        }
    }

    return true
}

//买卖股票的最佳时机 II 
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
//贪心算法

//只要后一天大于当前这一天的话，就可以把这个价差得到的利润给锁定下来
func maxProfit(prices []int) int {
    total := 0
    
    for i := 0; i < len(prices) - 1; i++ {
        if prices[i + 1] > prices[i] {
            total += prices[i + 1] - prices[i]
        }
    }
    return total
}



// 分发饼干
// 贪心算法
// 胃的大小 小于或等于 饼干的大小的话， 说明满足了一个孩子

// 时间、空间复杂度都是 O(n)
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    i, j := 0, 0
    for i < len(g) && j < len(s) {
        if g[i] <= s[j] {
            i++
        }
        j++
    }
    return i
}

// 模拟行走机器人
//疑惑：

// 欧式举例： 3^2 + 4 ^2 = 5 ^2, 如果坐标是3,4 ，那么结果是5，暂后在计算评分

// 方向： 上、右、下、左：0、1、2、3

func robotSim(commands []int, obstacles [][]int) int {
	result, curDir, mObstacles := float64(0), 0, make(map[string]bool)
	x, y, stepX, stepY := 0, 0, []int{0, 1, 0, -1}, []int{1, 0, -1, 0}

	for _, v := range obstacles {
		mObstacles[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}

	for _, v := range commands {
	   switch v {
	   case -1:
		   curDir = (curDir + 1 ) % 4
	   case -2:
		   curDir = (curDir + 3 ) % 4
	   default:
		   for i := 0; i < v; i++ {
			   tmpX, tmpY := x + stepX[curDir], y + stepY[curDir]

			   if mObstacles[fmt.Sprintf("%d-%d", tmpX, tmpY)] {
				   break
			   }

			   x, y = tmpX, tmpY
			   result = math.Max(float64(x * x + y * y), result)
		   }
	   }
	}

	return int(result)
}