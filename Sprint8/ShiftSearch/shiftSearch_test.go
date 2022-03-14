package shiftSearch_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	shiftSearch "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/ShiftSearch"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestShiftSearch(t *testing.T) {
	tests := []Test{{strings.NewReader(`9
3 9 1 2 5 10 9 1 7
2
4 10
`), "1 8 "}, {strings.NewReader(`5
1 2 3 4 5
3
10 11 12
`), "1 2 3 "}, {strings.NewReader(`50
352 153 127 830 665 764 309 306 127 830 665 764 309 584 478 42 745 580 679 224 927 762 861 406 148 160 863 698 797 342 571 65 768 603 702 247 950 785 884 429 422 63 766 601 700 245 948 783 882 427
5
83 786 621 720 265`), "3 9 16 20 26 32 36 42 46 "}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			shiftSearch.ShiftSearch(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
