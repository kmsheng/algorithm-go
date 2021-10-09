package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
	m "github.com/kmsheng/algorithm-go/model"
)

func TestRmDups(t *testing.T) {
	head := m.SliceToLinkedList([]int{3, 2, 1, 2, 3, 1})
	res := q.RmDups(&head)
	expected := m.SliceToLinkedList([]int{3, 2, 1})
	assert.Equal(t, *res, expected)
}
