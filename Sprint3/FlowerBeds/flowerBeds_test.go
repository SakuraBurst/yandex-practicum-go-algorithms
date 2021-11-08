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

func TestFlowerBeds(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		flowerBeds.FlowerBeds(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}