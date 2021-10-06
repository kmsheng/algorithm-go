package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestZeroMatrix(t *testing.T) {
	assert.Equal(
		t,
		q.ZeroMatrix(
			[][]int{
				{ 1, 0 },
				{ 3, 4 },
			},
		),
		[][]int{
			{ 0, 0 },
			{ 3, 0 },
		},
	)

	assert.Equal(
		t,
		q.ZeroMatrix(
			[][]int{
				{ 1, 2, 3 },
				{ 4, 0, 6 },
				{ 7, 8, 9 },
			},
		),
		[][]int{
			{ 1, 0, 3 },
			{ 0, 0, 0 },
			{ 7, 0, 9 },
		},
	)
}
