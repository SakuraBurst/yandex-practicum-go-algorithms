package allaOnTheAlgos_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	allaOnTheAlgos "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/AllaOnTheAlgos"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestAllaOnTheAlgos(t *testing.T) {
	tests := []Test{{strings.NewReader(`130
4
10 3 40 1
`), "4"}, {strings.NewReader(`100
2
7 5
`), "16"}, {strings.NewReader(`1
1
1
`), "1"}, {strings.NewReader(`20
3
6 4 10
`), "2"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			allaOnTheAlgos.AllaOnTheAlgos(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
