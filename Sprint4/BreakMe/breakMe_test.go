package breakMe_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	breakMe "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/BreakMe"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBreakMe(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		breakMe.BreakMe(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}