/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=20003
 *
 * [2] 两数相加
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    /*
       关键信息：
           1. 俩链表，节点放整数0-9
           2. 每个链表的所有节点构成一个整数,不会以0开头
           3. 求俩链表的对应整数的和，输出链表
           4. 链表是单链表
           5. 存在进位的情况
       分析：
          1. 因为有进位情况，需要从最尾部开始计算
           进位后，返回进位值，同时需要维护链表关系

          2. 链表长度不一，需要尾部对齐
            怎么对齐？链表反转？

            1.先进行链表反转
            2.链表头部加和，余数保留在当前节点，进位传入下一个节点
               递归函数
                 输入中有进位值,俩个链表
                 输出单链表
            3.下一个节点加和，加进位
            4.再次反转

           递归实现
    */

    //    l1 = reverse(l1)
    //    l2 = reverse(l2)

    val := l1.Val + l2.Val

    div := val / 10
    val = val % 10

    node := &ListNode{
        Val: val,
    }

    recure(node, div, l1.Next, l2.Next)
    //    return reverse(node)
    return node
}

func recure(pre *ListNode, val int, l1 *ListNode, l2 *ListNode) {

    if l1 == nil && l2 == nil {
        if val == 0 {
            return
        }
        node := &ListNode{
            Val: val,
        }
        pre.Next = node
        return
    }

    if l1 != nil && l2 == nil {
        val = l1.Val + val
        div := val / 10
        val = val % 10

        node := &ListNode{
            Val: val,
        }
        pre.Next = node
        recure(node, div, l1.Next, nil)
        return
    }

    if l1 == nil && l2 != nil {
        val = l2.Val + val
        div := val / 10
        val = val % 10

        node := &ListNode{
            Val: val,
        }
        pre.Next = node
        recure(node, div, nil, l2.Next)
        return
    }

    val = l1.Val + l2.Val + val
    div := val / 10
    val = val % 10

    node := &ListNode{
        Val: val,
    }
    pre.Next = node
    recure(node, div, l1.Next, l2.Next)

    return
}

func reverse(l *ListNode) *ListNode {

    var pre *ListNode = nil
    var cur *ListNode = l
    for {
        if cur.Next == nil {
            cur.Next = pre
            return cur
        }
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }

}

/*
总结：
    审题不清楚，思路正确，根本不需要反转
*/

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/

