package sameSums_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	sameSums "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7FinalEx/SameSums"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSameSums(t *testing.T) {
	tests := []Test{{strings.NewReader(`4
1 5 7 1
`), "True"}, {strings.NewReader(`3
2 10 9
`), "False"}, {strings.NewReader(`4
2 2 2 2
`), "True"}, {strings.NewReader(`6
7 9 3 4 6 7`), "True"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			sameSums.SameSums(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
