package model

import (
	"fmt"
	"strconv"
)

type LinkedNode struct {
	Value int
	Next *LinkedNode
}

type Fn func(*LinkedNode)

func SliceToLinkedList(slice []int) LinkedNode {
	head := LinkedNode{ Value: -1 }
	node := &head
	for _, v := range slice {
		node.Next = &LinkedNode{ Value: v }
		node = node.Next
	}
	return *head.Next
}

func Each(head *LinkedNode, fn Fn) {
	node := head
	for node != nil {
		fn(node)
		node = node.Next
	}
}

func Log(head *LinkedNode) {
	str := ""
	node := head
	for node != nil {
		str += strconv.Itoa(node.Value)
		if node.Next != nil {
			str += " -> "
		}
		node = node.Next
	}
	fmt.Println(str)
}
