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
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		shiftSearch.ShiftSearch(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
