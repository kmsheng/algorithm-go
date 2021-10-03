package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestIsOneAway(t *testing.T) {
	assert.True(t, q.IsOneAway("pale, ple"))
	assert.True(t, q.IsOneAway("pales, pale"))
	assert.True(t, q.IsOneAway("pale, bale"))
	assert.False(t, q.IsOneAway("pale, bake"))
}
