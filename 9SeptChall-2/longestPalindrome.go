package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	longest := ""

	for i := 0; i < len(s); i++ {
		// check for odd len
		oddPal := expandedAroundCentre(s, i, i)
		if len(oddPal) > len(longest) {
			longest = oddPal
		}
		evenPal := expandedAroundCentre(s, i, i+1)
		if len(evenPal) > len(longest) {
			longest = evenPal
		}
	}
	return longest
}

func expandedAroundCentre(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
