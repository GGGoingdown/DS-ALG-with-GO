package twopointers

import (
	"sort"
)

func FindSumOfThree(nums []int, target int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}

	copyNums := make([]int, length)
	copy(copyNums, nums)
	// sort by asc
	sort.Ints(copyNums)
	// using two pointers
	left, right := 0, length-1
	for left < right {
		j := left + 1
		for j < right {
			sum := copyNums[left] + copyNums[j] + copyNums[right]
			switch {
			case sum > target:
				right--
			case sum < target:
				j++
			case sum == target:
				return true
			}
		}
		left++
	}

	return false
}
