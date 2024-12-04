
/*
 * @lc app=leetcode.cn id=355 lang=golang
 * @lcpr version=20004
 *
 * [355] 设计推特
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type Twitter struct {
	users     map[int][]int
	followers map[int]map[int]struct{}
}

func Constructor() Twitter {
	/*

		实体:
		    用户: 关注，取消关注, 创建推文, 获取推文

		    推文:
			    1.  按时间顺序排序,
				2.  关注后，2发推文时，1的列表中有2的信息, 取消后没有
				      1和2的合并
		        3.  推文只追加

		设计：
		    推文设计为数组，只追加不删除。
			多个用户使用map存储map[int][]int, 每个用户的推文数组
			关注关系，设计为map[int]map[int]struct{}, 用户id关注了哪些用户
	*/

	return Twitter{
		users:     make(map[int][]int),
		followers: make(map[int]map[int]struct{}),
	}

}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if this.users[userId] == nil {
		this.users[userId] = []int{}
	}
	this.users[userId] = append(this.users[userId], tweetId)
}

func (this *Twitter) GetNewsFeed(userId int) []int {

	if this.followers[userId] == nil {
		this.followers[userId] = make(map[int]struct{})
		if _, ok := this.users[userId]; !ok {
			this.users[userId] = []int{}
		}
		return this.users[userId]
	}

	followers := this.followers[userId]

	// 合并k个有序数组, 并且只要最近的10条

	// 倒序，只取合并后的10条

	l1 := this.users[userId]
	tweets := make([][]int, len(followers)+1)

	tweets = append(tweets, l1)

	for userId := range followers {
		tweets = append(tweets, this.users[userId])
	}

	ans := getLastTenTweets(tweets)

	// sort.Slice(ans, func(i, j int) bool {
	// 	return ans[i] > ans[j]
	// })

	if len(ans) >= 10 {
		return ans[:10]
	}
	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {

	if this.followers[followerId] == nil {
		this.followers[followerId] = make(map[int]struct{})
	}
	this.followers[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followers[followerId] == nil {
		this.followers[followerId] = make(map[int]struct{})
	}
	delete(this.followers[followerId], followeeId)
}

func getLastTenTweets(ts [][]int) []int {
	var ans []int
	var newts = [][]int{}

	var maxTweet = -1
	var maxIndex = 0

	for i, tt := range ts {
		if len(tt) == 0 {
			continue
		}

		if maxTweet < tt[len(tt)-1] {
			maxTweet = tt[len(tt)-1]
			maxIndex = i
		}
	}

	if maxTweet == -1 {
		return ans
	}

	for i, tt := range ts {
		if i == maxIndex {
			newTt := tt[:len(tt)-1]
			if len(newTt) != 0 {
				newts = append(newts, newTt)
			}
		} else {
			newts = append(newts, tt)
		}
	}
	res := getLastTenTweets(newts)
	ans = append(ans, maxTweet)
	ans = append(ans, res...)
	return ans
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end



