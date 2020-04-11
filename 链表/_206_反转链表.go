package 链表

//link https://leetcode-cn.com/problems/reverse-linked-list/
//title 反转一个单链表

/* 迭代方法
* 1 -> 2 -> 3 -> 4 -> null
* null <- 1 <- 2 <- 3 <- 4
 */
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}
