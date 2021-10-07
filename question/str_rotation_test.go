package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestStrRotation(t *testing.T) {
	assert.True(t, q.StrRotation("waterbottle", "erbottlewat"))
	assert.True(t, q.StrRotation("waterbottle", "rbottlewate"))
	assert.True(t, q.StrRotation("waterbottle", "lewaterbott"))
}
