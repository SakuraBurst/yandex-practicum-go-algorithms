package searchTree_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	searchTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/SearchTree"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSearchTree(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		searchTree.SearchTree(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}