package balancedTree_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	balancedTree "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/BalancedTree"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBalancedTree(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		balancedTree.BalancedTree(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}