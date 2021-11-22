package competition_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	competition "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/Competition"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCompetition(t *testing.T) {
	tests := []Test{{strings.NewReader(`2
0 1`), "2"}, {strings.NewReader(`2
0 1 0`), "2"}, {strings.NewReader(`10
0 0 1 0 0 0 1 0 0 1`), "4"}, {strings.NewReader(`72
1 1 1 1 0 1 1 0 0 0 0 0 1 1 0 0 1 1 0 0 1 0 0 1 0 0 1 1 0 0 1 0 0 0 0 1 1 1 0 0 0 1 0 1 1 1 0 1 1 0 0 1 0 0 0 1 0 1 0 0 0 1 1 1 1 1 1 0 1 0 0 0
`), "46"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		competition.Competition(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
