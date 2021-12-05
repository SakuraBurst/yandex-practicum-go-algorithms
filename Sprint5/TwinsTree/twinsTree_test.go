package twinsTree_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	twinsTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/TwinsTree"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestTwinsTree(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		twinsTree.TwinsTree(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}