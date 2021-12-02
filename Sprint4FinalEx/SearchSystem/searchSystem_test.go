package searchSystem_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	searchSystem "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4FinalEx/SearchSystem"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSearchSystem(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		searchSystem.SearchSystem(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}