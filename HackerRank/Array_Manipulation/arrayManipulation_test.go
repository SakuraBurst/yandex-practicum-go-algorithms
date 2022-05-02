package Array_Manipulation_test

import (
	"bytes"
	arraymanipulation "github.com/SakuraBurst/yandex-practicum-go-algorithms/HackerRank/Array_Manipulation"
	"io"
	"strings"
	"testing"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBorderControl(t *testing.T) {
	tests := []Test{{strings.NewReader(`5 3
1 2 100
2 5 100
3 4 100`), "200\n"}, {strings.NewReader(`10 3
1 5 3
4 8 7
6 9 1`), "10\n"}, {strings.NewReader(`10 4
2 6 8
3 5 7
1 8 1
5 9 15`), "31\n"},
	}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			arraymanipulation.PrepareData(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
