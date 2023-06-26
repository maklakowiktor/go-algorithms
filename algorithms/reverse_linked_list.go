package algorithms

type ListNode struct {
	Next *ListNode
	Val int
}

// 1 -> 2 -> 3 -> 4
// 1 -> nil
// 2 -> 1
// 3 -> 2
// 4 -> 3
// return 4 -> 3 -> 2 -> 1
func (head *ListNode) Reverse() *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
