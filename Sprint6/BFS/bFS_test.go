package bFS_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	bFS "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/BFS"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBFS(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		bFS.BFS(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}