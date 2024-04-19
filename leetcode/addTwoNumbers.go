/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp int = 0
	node := &ListNode{}
	first := node
	for {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//log.Println(node.Val)
		node.Val = (n1 + n2 + tmp) % 10
		tmp = (n1 + n2 + tmp) / 10
		if l1 == nil && l2 == nil {
			break
		}
		newNode := &ListNode{Val: 0, Next: nil}
		node.Next = newNode
		node = node.Next
	}
	if tmp != 0 {
		node.Next = &ListNode{
			Val: tmp,
		}
	}
	return first
}
