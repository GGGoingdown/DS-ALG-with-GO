package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	/*
		"kayak"
		"hello"
		"RACEACAR"
		"A"
		"ABCDABCD"
		"DCBAABCD"
		"ABCBA"
	*/

	testcase := []struct {
		input    string
		expected bool
	}{
		{"kayak", true},
		{"hello", false},
		{"RACEACAR", false},
		{"A", true},
	}
	for _, e := range testcase {
		// if result := isPalindrome(e.input); result != e.expected {
		// 	t.Errorf("'%s' expected %v", e.input, e.expected)
		// }
		assert.EqualValuesf(t, e.expected, isPalindrome(e.input), "%s is %v", e.input, e.expected)

	}
}

func TestIsPalindromeII(t *testing.T) {
	/*
		"kayak"
		"hello"
		"RACEACAR"
		"A"
		"ABCDABCD"
		"DCBAABCD"
		"ABCBA"
	*/

	testcase := []struct {
		input    string
		expected bool
	}{
		{"madame", true},
		{"dead", true},
		{"abca", true},
		{"tebbem", false},
		{"eeccccbebaeeabebccceea", false},
		{"ognfjhgbjhzkqhzadmgqbwqsktzqwjexqvzjsopolnmvnymbbzoofzbbmynvmnloposjzvqxejwqztksqwbqgmdazhqkzhjbghjfno", false},
	}
	for _, e := range testcase {
		assert.EqualValuesf(t, e.expected, IsPalindromeII(e.input), "%s is %v", e.input, e.expected)
	}
}
