package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	// test case 1
	input := "babad"
	expected := "bab"
	result := longestPalindrome(input)
	if result != expected && result != "aba" {
		t.Errorf("Test failed for input %s, got %s, expected %s or %s", input, result, expected, "aba")
	}
	// test case 2
	input = "cbbd"
	expected = "bb"
	result = longestPalindrome(input)
	if result != expected {
		t.Errorf("Test failed for input %s, got %s, expected %s or %s", input, result, expected, "aba")
	}
	// test case 3
	input = "a"
	expected = "a"
	result = longestPalindrome(input)
	if result != expected {
		t.Errorf("Test failed for input %s, got %s, expected %s or %s", input, result, expected, "aba")
	}
	// test case 4
	input = "ac"
	expected = "a"
	result = longestPalindrome(input)
	if result != expected {
		t.Errorf("Test failed for input %s, got %s, expected %s", input, result, expected)
	}
}
