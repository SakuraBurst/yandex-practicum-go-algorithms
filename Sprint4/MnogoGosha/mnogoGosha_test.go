package mnogoGosha_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	mnogoGosha "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/MnogoGosha"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestMnogoGosha(t *testing.T) {
	tests := []Test{{strings.NewReader(`10 2
gggggooooogggggoooooogggggssshaa`), "0 5 "}, {strings.NewReader(`3 4
allallallallalla`), "0 1 2 "}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		mnogoGosha.MnogoGosha(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
