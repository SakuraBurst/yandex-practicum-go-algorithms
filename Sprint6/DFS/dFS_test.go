package dFS_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	dFS "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/DFS"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestDFS(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		dFS.DFS(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}