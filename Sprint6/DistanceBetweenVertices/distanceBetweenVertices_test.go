package distanceBetweenVertices_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	distanceBetweenVertices "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/DistanceBetweenVertices"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestDistanceBetweenVertices(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		distanceBetweenVertices.DistanceBetweenVertices(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}