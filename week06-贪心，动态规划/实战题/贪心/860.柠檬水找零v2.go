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

	money := []int{}
	ans := false
	for i := range bills {
		if bills[i] == 5 {
			money = append(money, 5)
			ans = true
			continue
		}

		if len(money) == 0 {
			ans = false
			break
		}

		bill := bills[i]
		for j := 0; j < len(money); j++ {
			bill = bill - money[j]
			if bill == 0 {
				ans = true
				// money应该去掉j对应的值
				money = money[j:]
				money = append(money, bills[i])
				break
			}
		}

		if bill > 0 {
			ans = false
			break
		}
	}
	return ans
}

/*
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

