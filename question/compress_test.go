package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestCompress(t *testing.T) {
	assert.Equal(t, "a2b1c5a3", q.Compress("aabcccccaaa"))
	assert.Equal(t, q.Compress("wtf"), "wtf")
	assert.Equal(t, q.Compress("wtffffff"), "w1t1f6")
}
