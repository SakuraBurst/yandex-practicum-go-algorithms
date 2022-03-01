package frequentWord_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	frequentWord "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/FrequentWord"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFrequentWord(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		frequentWord.FrequentWord(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
