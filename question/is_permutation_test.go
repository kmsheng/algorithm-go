package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestIsPermutation(t *testing.T) {
	assert.True(t, q.IsPermutation("abc", "bca"))
	assert.False(t, q.IsPermutation("abd", "bca"))
}
