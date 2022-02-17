package levenshteinDistance_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	levenshteinDistance "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7FinalEx/LevenshteinDistance"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestLevenshteinDistance(t *testing.T) {
	tests := []Test{{strings.NewReader(`abacaba
abaabc
`), "2"}, {strings.NewReader(`innokentiy
innnokkentia
`), "3"}, {strings.NewReader(`r
x
`), "1"}, {strings.NewReader(`dxqrpmratn
jdpmykgmaitn`), "8"}, {strings.NewReader(`abcdefghj
mnabcdelghj`), "3"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			levenshteinDistance.LevenshteinDistance(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
