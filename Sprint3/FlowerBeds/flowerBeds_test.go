package flowerBeds_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	flowerBeds "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/FlowerBeds"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFlowerBeds(t *testing.T) {
	tests := []Test{{strings.NewReader(`4
7 8
7 8
2 3
6 10
`), `2 3
6 10
`}, {strings.NewReader(`4
2 3
5 6
3 4
3 4
`), `2 4
5 6
`}, {strings.NewReader(`6
1 3
3 5
4 6
5 6
2 4
7 10
`), `1 6
7 10
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		flowerBeds.FlowerBeds(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
