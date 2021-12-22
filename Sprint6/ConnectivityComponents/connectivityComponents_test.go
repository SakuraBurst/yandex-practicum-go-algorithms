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

func TestConnectivityComponents(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		connectivityComponents.ConnectivityComponents(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}