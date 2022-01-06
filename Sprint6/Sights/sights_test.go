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

func TestSights(t *testing.T) {
	tests := []Test{{strings.NewReader(`4 4
1 2 1
2 3 3
3 4 5
1 4 2
`), `0 1 4 2 
1 0 3 3 
4 3 0 5 
2 3 5 0 
`}, {strings.NewReader(`3 2
1 2 1
1 2 2
`), `0 1 -1 
1 0 -1 
-1 -1 0 
`}, {strings.NewReader(`2 0
`), `0 -1 
-1 0 
`}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			sights.Sights(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
