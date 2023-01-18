package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSumOfThree(t *testing.T) {
	testcase := []struct {
		nums     []int
		target   int
		expected bool
	}{
		{nums: []int{3, 7, 1, 2, 8, 4, 5}, target: 10, expected: true},
		{[]int{3, 7, 1, 2, 8, 4, 5}, 21, false},
		{[]int{3, 7, 1, 2, 8, 4, 5}, 20, true},
		{[]int{1, -1, 0}, 0, true},
	}

	for _, e := range testcase {
		assert.Equal(t, e.expected, FindSumOfThree(e.nums, e.target))
	}

}
