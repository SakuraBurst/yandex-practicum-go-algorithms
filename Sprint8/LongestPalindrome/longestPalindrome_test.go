package longestPalindrome_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	longestPalindrome "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/LongestPalindrome"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestLongestPalindrome(t *testing.T) {
	tests := []Test{
		{strings.NewReader("aaaabb\n"), "aabbaa"},
		{strings.NewReader("pabcd\n"), "a"},
		{strings.NewReader("aaabbb\n"), "ababa"},
		{strings.NewReader("paqokjqvzi\n"), "qaq"},
	}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			longestPalindrome.LongestPalindrome(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
