package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
	m "github.com/kmsheng/algorithm-go/model"
)

func TestRmNode(t *testing.T) {
	head := m.SliceToLinkedList([]int{1, 2, 3, 4})
	res := q.RmNode(&head)
	expected := m.SliceToLinkedList([]int{2, 3, 4})
	assert.Equal(t, *res, expected)
}
