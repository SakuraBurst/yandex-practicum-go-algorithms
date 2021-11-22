package mnogoGosha_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	mnogoGosha "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/MnogoGosha"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestMnogoGosha(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		mnogoGosha.MnogoGosha(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}