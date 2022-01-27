package horoscopes_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	horoscopes "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Horoscopes"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestHoroscopes(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		horoscopes.Horoscopes(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
