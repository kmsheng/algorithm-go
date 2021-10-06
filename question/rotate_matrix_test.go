package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestRotateMatrix(t *testing.T) {
	// 2x2
	assert.Equal(
		t,
		q.RotateMatrix(
			[][]int{
				{ 1, 2 },
				{ 3, 4 },
			},
		),
		[][]int{
			{ 3, 1 },
			{ 4, 2 },
		},
	)

	// 3x3
	assert.Equal(
		t,
		q.RotateMatrix(
			[][]int{
				{ 1, 2, 3 },
				{ 4, 5, 6 },
				{ 7, 8, 9 },
			},
		),
		[][]int{
			{ 7, 4, 1 },
			{ 8, 5, 2 },
			{ 9, 6, 3 },
		},
	)

	// 4x4
	assert.Equal(
		t,
		q.RotateMatrix(
			[][]int{
				{ 1, 2, 3, 4 },
				{ 5, 6, 7, 8 },
				{ 9, 10, 11, 12 },
				{ 13, 14, 15, 16 },
			},
		),
		[][]int{
				{ 13, 9, 5, 1 },
				{ 14, 10, 6, 2 },
				{ 15, 11, 7, 3 },
				{ 16, 12, 8, 4 },
		},
	)

	// 5x5
	assert.Equal(
		t,
		q.RotateMatrix(
			[][]int{
				{ 1, 2, 3, 4, 5 },
				{ 6, 7, 8, 9, 10 },
				{ 11, 12, 13, 14, 15 },
				{ 16, 17, 18, 19, 20 },
				{ 21, 22, 23, 24, 25 },
			},
		),
		[][]int{
			{ 21, 16, 11, 6, 1 },
			{ 22, 17, 12, 7, 2 },
			{ 23, 18, 13, 8, 3 },
			{ 24, 19, 14, 9, 4 },
			{ 25, 20, 15, 10, 5 },
		},
	)
}
