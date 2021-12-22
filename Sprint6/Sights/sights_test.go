package sights_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	sights "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/Sights"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSights(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		sights.Sights(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}