package twopointers

import (
	"strings"
)

func ReverseWords(sentence string) string {
	tmp_sentence := strings.Split(sentence, " ")

	s_sentence := removeSpace(tmp_sentence)

	if len(s_sentence) <= 1 {
		return strings.Join(s_sentence, "")
	}

	left, right := 0, len(s_sentence)-1
	for left < right {
		s_sentence[left], s_sentence[right] = s_sentence[right], s_sentence[left]
		left++
		right--
	}

	return strings.Join(s_sentence, " ")
}

func removeSpace(str []string) []string {
	newStr := make([]string, 0, len(str))

	for _, s := range str {
		if s != "" {
			newStr = append(newStr, s)
		}
	}

	return newStr
}
