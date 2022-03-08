package borderControl_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	borderControl "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/BorderControl"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBorderControl(t *testing.T) {
	tests := []Test{{strings.NewReader(`abcdefg
abdefg
`), "OK"}, {strings.NewReader(`helo
hello
`), "OK"}, {strings.NewReader(`dog
fog
`), "OK"}, {strings.NewReader(`mama
papa
`), "FAIL"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			borderControl.BorderControl(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
