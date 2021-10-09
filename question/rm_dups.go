package question

import (
	m "github.com/kmsheng/algorithm-go/model"
)

func RmDups(head *m.LinkedNode) *m.LinkedNode {
	m.Each(head, func(node *m.LinkedNode) {
		prev := node
		m.Each(node.Next, func(n2 *m.LinkedNode) {
			if node.Value == n2.Value {
				prev.Next = n2.Next
			} else {
				prev = n2
			}
		})
	})
	return head
}
