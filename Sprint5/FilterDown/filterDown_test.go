package filterDown_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	filterDown "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/FilterDown"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFilterDown(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		filterDown.FilterDown(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}