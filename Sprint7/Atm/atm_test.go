package atm_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	atm "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Atm"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestAtm(t *testing.T) {
	tests := []Test{{strings.NewReader(`5
3
3 2 1`), "5"}, {strings.NewReader(`3
2
2 1`), "2"}, {strings.NewReader(`8
1
5`), "0"}, {strings.NewReader(`9
7
10 2 3 7 6 5 9`), "6"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			atm.Atm(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
