package twopointers

func isPalindrome(inputString string) bool {
	length := len(inputString)
	// if string length less than equal 1, is palidrome
	if length <= 1 {
		return true
	}

	left, right := 0, length-1
	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func IsPalindromeII(inputString string) bool {
	length := len(inputString)
	// if string length less than equal 1, is palidrome
	if length <= 1 {
		return true
	}
	// 有一次的機會可以bypass
	counter := 1
	// two pointers
	left, right := 0, length-1
	for left < right {
		// if string not the same
		if inputString[left] != inputString[right] {
			// check counter is greater than 0
			if counter > 0 {
				switch {
				case inputString[left+1] == inputString[right]:
					left++
				case inputString[left] == inputString[right-1]:
					right--
				default:
					return false
				}
				counter--
			} else {
				return false
			}
		}
		left++
		right--
	}

	return true

}
