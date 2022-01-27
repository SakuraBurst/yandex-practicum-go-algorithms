package journey_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	journey "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Journey"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestJourney(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		journey.Journey(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
