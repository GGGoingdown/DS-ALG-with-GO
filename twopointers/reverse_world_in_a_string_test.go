package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	testcase := []struct {
		input    string
		expected string
	}{
		{"We love Go", "Go love We"},
		{"Hello     World", "World Hello"},
	}

	for _, e := range testcase {
		assert.Equal(t, e.expected, ReverseWords(e.input))
	}

}
