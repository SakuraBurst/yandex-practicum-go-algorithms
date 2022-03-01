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
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		longestPalindrome.LongestPalindrome(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
