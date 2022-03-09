package compareTwoStrings_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	compareTwoStrings "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/CompareTwoStrings"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCompareTwoStrings(t *testing.T) {
	tests := []Test{{strings.NewReader(`gggggbbb
bbef
`), "-1"}, {strings.NewReader(`z
aaaaaaa
`), "1"}, {strings.NewReader(`ccccz
aaaaaz
`), "0"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			compareTwoStrings.CompareTwoStrings(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
