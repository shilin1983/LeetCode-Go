package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

// * 整数数组转链表
func Int2List(nums []int) *ListNode {
	var (
		length int       = len(nums)
		head   *ListNode = nil
	)

	if length != 0 {
		for i := length - 1; i >= 0; i-- {
			node := &ListNode{Val: nums[i]}
			node.Next = head
			head = node
		}
	}

	return head
}
