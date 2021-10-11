package question

import (
	m "github.com/kmsheng/algorithm-go/model"
	"reflect"
)


func RmNode(node *m.LinkedNode) *m.LinkedNode {
	target := node
	for target.Next != nil {
		target.Value = target.Next.Value
		if (! reflect.ValueOf(target.Next).IsNil()) && reflect.ValueOf(target.Next.Next).IsNil() {
			target.Next = nil
			break
		}
		target = target.Next
	}
	return node
}
