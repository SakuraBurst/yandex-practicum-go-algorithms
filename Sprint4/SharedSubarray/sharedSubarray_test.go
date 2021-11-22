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

func TestSharedSubarray(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		sharedSubarray.SharedSubarray(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}