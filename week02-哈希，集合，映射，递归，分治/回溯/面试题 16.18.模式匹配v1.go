
/*
 * @lc app=leetcode.cn id=面试题 16.18 lang=golang
 * @lcpr version=20003
 *
 * [面试题 16.18] 模式匹配
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func patternMatching(pattern string, value string) bool {
	/*

		pattern 由a,b构成, 表示不同的字符串,  1,1000
		   a,b可以为""

		输入value， 判断是否符合pattern     0,1000

		子问题：
		   1. 先判断pattern第一个字符，对应什么模式
		       当value中 a对应第一个字母时, 将pattern中的a进行替换

			    确定b对应什么模式 ==> 递归
				   调整value，进入下一次判断

		   2. 当value为空时，返回true
		      pattern中没有时，但value还有，返回false
			  pattern中有，但value已经为空， 返回false

		   3. 如果失败得回到上一层


		   a()b()
	*/

	a := []byte("")

	for i := 0; i < len(value); i++ {
		a = append(a, value[i])
		if r1(i+1, len(pattern), string(a), value) {
			return true
		}
	}

	return false
}

func r1(i int, n int, a string, value string) bool {
	b := []byte("")
	for j := i; j < len(value); j++ {
		b = append(b, value[i])
		l := r(n, a, string(b))
		for _, v := range l {
			if v == value {
				return true
			}
		}
	}
	return false
}

func r(n int, a, b string) []string {
	if n == 0 {
		return []string{""}
	}
	var res = []string{}

	for k := 1; k <= n; k++ {
		al := r(n-k, a, b)
		bl := r(k-1, a, b)

		for _, v1 := range al {
			for _, v2 := range bl {
				v := fmt.Sprintf("%s%s%s%s", a, v1, b, v2)
				res = append(res, v)
			}
		}
	}
	return res
}

/*

总结：老感觉差一点，可以做出, 但实际还是想的有问题

   最终思路没错，细节与程序编码有问题，时间消耗估计1个多小时

   这种方式不对，不能这么做算法题了

*/

// @lc code=end

/*
// @lcpr case=start
// "abba"\n"dogcatcatdog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dogcatcatfish"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n"dogcatcatdog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dogdogdogdog"\n
// @lcpr case=end

*/

