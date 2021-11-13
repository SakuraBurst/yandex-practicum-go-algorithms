package housebuying_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	housebuying "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/HouseBuying"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestHousebuying(t *testing.T) {
	tests := []Test{{strings.NewReader(`3 300
999 999 999`), "0"}, {strings.NewReader(`3 1000
350 999 200`), "2"}, {strings.NewReader(`1 4
3`), "1"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		housebuying.HouseBuying(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
