package anagramGrouping_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	anagramGrouping "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/AnagramGrouping"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestAnagramGrouping(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		anagramGrouping.AnagramGrouping(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}