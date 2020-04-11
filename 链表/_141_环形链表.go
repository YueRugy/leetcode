package 链表

//link https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	for quick, slow := head, head.Next; quick.Next != nil; quick, slow = quick.Next.Next, slow.Next {
		if quick = slow {
			return true
		}
	}
	return false
}
