import 链表
// link https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
//title 删除链表中的节点
 type ListNode struct {
      Val int
      Next *ListNode
  }
func deleteNode(node *ListNode) {
    node.Val=node.Next.Val
    node.Next=node.Next.Next
}
