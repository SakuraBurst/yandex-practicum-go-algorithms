package insertingRows_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	insertingRows "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/InsertingRows"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestInsertingRows(t *testing.T) {
	tests := []Test{{strings.NewReader(`abacaba
3
queue 2
deque 0
stack 7
`), "dequeabqueueacabastack"}, {strings.NewReader(`kukareku
2
p 1
q 2
`), "kpuqkareku"}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		insertingRows.InsertingRows(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
