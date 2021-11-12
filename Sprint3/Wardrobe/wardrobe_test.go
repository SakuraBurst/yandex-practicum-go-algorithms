package wardrobe_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	wardrobe "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/Wardrobe"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestWardrobe(t *testing.T) {
	tests := []Test{{strings.NewReader(`7
0 2 1 2 0 0 1`), "0 0 0 1 1 2 2"}, {strings.NewReader(`5
2 1 2 0 1`), "0 1 1 2 2"}, {strings.NewReader(`6
2 1 1 2 0 2`), "0 1 1 2 2 2"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		wardrobe.Wardrobe(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
