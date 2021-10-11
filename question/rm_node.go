package question

import (
	m "github.com/kmsheng/algorithm-go/model"
	u "github.com/kmsheng/algorithm-go/util"
)


func RmNode(node *m.LinkedNode) *m.LinkedNode {
	target := node
	for target.Next != nil {
		target.Value = target.Next.Value
		if (! u.IsNil(target.Next)) && u.IsNil(target.Next.Next) {
			target.Next = nil
			break
		}
		target = target.Next
	}
	return node
}
