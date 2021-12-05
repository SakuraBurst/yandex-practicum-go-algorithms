package treeAnagramm_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	treeAnagramm "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/TreeAnagramm"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestTreeAnagramm(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		treeAnagramm.TreeAnagramm(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}