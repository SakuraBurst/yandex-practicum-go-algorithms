package railroads_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	railroads "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6FinalEx/Railroads"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestRailroads(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		railroads.Railroads(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
