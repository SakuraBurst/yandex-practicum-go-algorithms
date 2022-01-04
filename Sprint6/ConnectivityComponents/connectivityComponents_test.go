package connectivityComponents_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	connectivityComponents "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/ConnectivityComponents"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestConnectivityComponents(t *testing.T) {
	tests := []Test{{strings.NewReader(`6 3
1 2
6 5
2 3
`), `3
1 2 3 
4 
5 6 
`}, {strings.NewReader(`2 0
`), `2
1 
2 
`}, {strings.NewReader(`4 3
2 3
2 1
4 3
`), `1
1 2 3 4 
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		connectivityComponents.ConnectivityComponents(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
