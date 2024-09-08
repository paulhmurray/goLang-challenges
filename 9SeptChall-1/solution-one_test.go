package main

import "testing"

func TestSolveProblem(t *testing.T) {
	// arrange
	cases := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 4},
		{input: 3, expected: 9},
	}
	// iterate over each test case
	for _, c := range cases {
		// Act: Call the function being tested
		result := solveProblem(c.input)
		// Assert: check if the result matches the expected output
		if result != c.expected {
			t.Errorf("solveProblem(%d) == %d", c.input, c.expected)
		}
	}
}
