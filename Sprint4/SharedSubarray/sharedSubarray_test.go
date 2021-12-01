package sharedSubarray_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	sharedSubarray "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/SharedSubarray"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSharedSubarray(t *testing.T) {
	tests := []Test{{strings.NewReader(`5
1 2 3 2 1
5
3 2 1 5 6
`), "3"}, {strings.NewReader(`5
1 2 3 4 5
3
4 5 9
`), "2"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		sharedSubarray.SharedSubarray(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
