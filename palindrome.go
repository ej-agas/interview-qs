package interview_qs

import "strings"

func IsPalindrome(s string) bool {
	// transform to lowercase the string and remove all spaces
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")

	for i := 0; i < len(s)/2; i++ { // run loop until the middle part of the string
		left := s[i]           // left side of the string
		right := s[len(s)-1-i] // right side of the string

		if left != right {
			return false
		}
	}

	return true
}
