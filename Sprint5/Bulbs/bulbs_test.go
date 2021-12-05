package bulbs_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	bulbs "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5/Bulbs"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBulbs(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		bulbs.Bulbs(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}