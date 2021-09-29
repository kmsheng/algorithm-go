package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestUrlify(t *testing.T) {
	assert.Equal(t, q.Urlify("  B C  "), "%20%20B%20C%20%20")
	assert.Equal(t, q.Urlify("ABC"), "ABC")
}
