package question

import (
	m "github.com/kmsheng/algorithm-go/model"
)

func Kth(head *m.LinkedNode, k int) *m.LinkedNode {
	node := head
	i := 0
	for node != nil {
		if i == k {
			return node
		}
		node = node.Next
		i++
	}
	return nil
}
