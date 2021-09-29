package question_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	q "github.com/kmsheng/algorithm-go/question"
)

func TestIsPalindromePermutation(t *testing.T) {

	bool1, permutations := q.IsPalindromePermutation("tacocatxyz")
	assert.False(t, bool1)

	bool2, permutations := q.IsPalindromePermutation("tacocat")
	assert.True(t, bool2)

	res := []string{"catotac", "actotca", "atcocta", "ctaoatc", "tcaoact", "tacocat"}
	assert.Equal(t, permutations, res)
}
