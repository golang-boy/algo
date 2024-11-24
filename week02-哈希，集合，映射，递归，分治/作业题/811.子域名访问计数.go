/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	/*
		   输入： 字符串列表,每个字符串形式如： rep  d1.d2.d3

		   输出：输出一个列表,每个字符为输入中所有域名的子域名访问的rep 形式

		   边界：
		       输入长度限制在1,100
			   子域名长度限制在1,100
			   rep的范围1,10000

		   思路：
		     1. 从输入中找到所有子域名
			 2. 累积子域名对应的次数
	*/

	stats := map[string]int{}
	ans := []string{}

	for _, cpdomain := range cpdomains {
		domains := GetSubDomains(cpdomain)
		for name, rep := range domains {
			stats[name] += rep
		}
	}

	for name, rep := range stats {
		ans = append(ans, fmt.Sprintf("%d %s", rep, name))
	}

	return ans
}

type Domain struct {
	Rep  int
	Name string
}

func GetSubDomains(cpdomain string) map[string]int {
	list := strings.Split(cpdomain, " ")
	rep, _ := strconv.Atoi(list[0])

	list = strings.Split(list[1], ".")

	res := map[string]int{}
	pre := ""
	for i := len(list) - 1; i >= 0; i-- {
		name := list[i]
		if pre != "" {
			name = list[i] + "." + pre
		}
		res[name] = rep
		pre = name
	}
	return res
}

/*
总结：
  使用map去统计，时间复杂度：O(n*m), n表示输入的数组的长度，m表示子域名的数量
*/

// @lc code=end

