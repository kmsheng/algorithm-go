package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestIsUniq(t *testing.T) {
	assert.True(t, q.IsUniq("abc"))
	assert.False(t, q.IsUniq("cbc"))
}
