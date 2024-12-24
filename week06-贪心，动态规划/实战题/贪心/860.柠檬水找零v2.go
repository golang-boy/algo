/*
 * @lc app=leetcode.cn id=860 lang=golang
 * @lcpr version=20004
 *
 * [860] 柠檬水找零
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lemonadeChange(bills []int) bool {

	/*
	  尝试使用搜索解
	   判断能不能正确找零
	    bills[i] in [5, 10, 20]
	    len(bills) [1, 100000]

	    手头的钱money是状态
	      money
	     什么时候找钱，什么时候不找钱？
	      if bills[i] == 5
	         不找钱
	         money = append(money, bills[i])
	      else
	        for j:=0;j<len(money);j++{
	           if bills[i] - money[j] == 0 {
	              结束
	           } else if bills[i] - money[j] > 0 {
	              递归
	           }
	        }
	*/

	money := map[int]int{}

	for i := range bills {
		if !exchange(bills[i]-5, money) { // 不能找零，则结束
			return false
		}
		money[bills[i]]++
	}
	return true
}

func exchange(amount int, money map[int]int) bool {
	// 方法一

	// if amount == 0 {
	// 	return true
	// }
	// for _, v := range []int{20, 10, 5} {
	// 	nums := money[v]
	// 	for z := 1; z <= nums; z++ {
	// 		if amount-v == 0 {
	// 			money[v]--
	// 			return true
	// 		} else if amount-v > 0 {
	// 			amount -= v
	// 			money[v]--
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }
	// return false

	// 61/61 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 38.33 % of golang submissions (9.4 MB)

	// 方法二
	for _, v := range []int{20, 10, 5} {
		for amount >= v && money[v] > 0 {
			amount -= v
			money[v]--
		}
	}
	return amount == 0

	//	61/61 cases passed (0 ms)
	// Your runtime beats 100 % of golang submissions
	// Your memory usage beats 37.5 % of golang submissions (9.5 MB)
}

/*
    总结：
	    先判断最大的，每次尝试减去最大的面值，减了之后，判断是否继续，对应的money得减少

		需要注意的是，money[v] 等建一个副本，用于循环，否则某某会修改它的值，某些循环导致失败
*/

// @lc code=end

/*
// @lcpr case=start
// [5,5,5,10,20]\n
// @lcpr case=end

// @lcpr case=start
// [5,5,10,10,20]\n
// @lcpr case=end

*/

