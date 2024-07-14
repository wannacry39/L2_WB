package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {

	testTable := []struct {
		inputLine string
		expected  string
	}{
		{
			inputLine: "",
			expected:  "",
		},
		{
			inputLine: "a4bc2d5e",
			expected:  "aaaabccddddde",
		},
		{
			inputLine: "abcd",
			expected:  "abcd",
		},
		{
			inputLine: "qwe\\4\\5",
			expected:  "qwe45",
		},
		{
			inputLine: "qwe\\45",
			expected:  "qwe44444",
		},
	}

	for _, testcase := range testTable {
		result, _ := Unpack(testcase.inputLine)
		t.Logf("Calling Unpack(%v), result %s\n", testcase.inputLine, result)

		if result != testcase.expected {
			t.Errorf("Incorrect result. Expect %s, got %s", testcase.expected, result)
		}
	}
}
