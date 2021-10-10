package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
	m "github.com/kmsheng/algorithm-go/model"
)

func TestKth(t *testing.T) {
	head := m.SliceToLinkedList([]int{1, 2, 3, 4, 5})
	res := q.Kth(&head, 2)
	assert.Equal(t, res.Value, 3)

	res = q.Kth(&head, -1)
	assert.Nil(t, res)
}
